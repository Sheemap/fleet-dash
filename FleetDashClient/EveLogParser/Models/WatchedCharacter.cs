using System.Text;
using System.Text.RegularExpressions;
using EveLogParser.Constants;
using Serilog;
using YamlDotNet.Core;
using YamlDotNet.Serialization;

namespace EveLogParser.Models;

internal class WatchedCharacter
{
    public WatchedCharacter(string characterId, string? overviewPath)
    {
        CharacterId = characterId;
        OverviewPath = overviewPath;
    }

    public string CharacterId { get; }
    private string? OverviewPath { get; }


    private IEnumerable<KeyValuePair<string, string>>? _labelRegex;
    
    private IEnumerable<KeyValuePair<string, string>> LabelRegex => _labelRegex ??= GetLabelRegex();

    public string GetCharacterRegex(bool includeAlliance)
    {
        var charRegex = new StringBuilder();
        foreach (var (_, value) in LabelRegex.Where(kvp => includeAlliance || kvp.Key != "alliance"))
        {
            charRegex.Append(value);
        }

        return charRegex + EnglishRegex.Weapon;
    }

    private IEnumerable<KeyValuePair<string,string>> GetLabelRegex()
    {
        if (string.IsNullOrWhiteSpace(OverviewPath))
        {
            return null;
        }
        
        using var reader = new StreamReader(OverviewPath);
        var overviewYaml = reader.ReadToEnd();

        var deserializer = new Deserializer();
        var settings = deserializer.Deserialize<Dictionary<string, object>>(overviewYaml);

        try
        {
            var shipSettings = ExtractShipLabelSettings(settings);

            var labelRegexList = new List<KeyValuePair<string, string>>();
            foreach (var label in shipSettings)
            {
                // TODO: This spew of if statements is pretty ugly. Can we refactor?
                var labelRegex = new StringBuilder();
                var (_, typeValue) = label.FirstOrDefault(x => x.Key == "type");
                if (typeValue == "linebreak")
                {
                    var linebreakKvp = new KeyValuePair<string, string>(typeValue, " ");
                    labelRegexList.Add(linebreakKvp);
                    continue;
                }
                
                var (_, fontValue) = label.FirstOrDefault(x => x.Key == "fontsize");
                if (fontValue != null)
                {
                    labelRegex.Append($"<font size={fontValue}>");
                }

                var (_, colorValue) = label.FirstOrDefault(x => x.Key == "color");
                if (colorValue != null)
                {
                    labelRegex.Append("<color=.*?>");
                }

                var (_, preValue) = label.FirstOrDefault(x => x.Key == "pre");
                if (preValue != null)
                {
                    labelRegex.Append(Regex.Escape(preValue));
                }

                var (_, italicValue) = label.FirstOrDefault(x => x.Key == "italic");
                var itallicsUsed = IsOverviewBoolTruthy(italicValue);
                if (itallicsUsed)
                {
                    labelRegex.Append("<i>");
                }

                var (_, underlineValue) = label.FirstOrDefault(x => x.Key == "underline");
                var underlineUsed = IsOverviewBoolTruthy(underlineValue);
                if (underlineUsed)
                {
                    labelRegex.Append("<u>");
                }

                var (_, boldValue) = label.FirstOrDefault(x => x.Key == "bold");
                var boldUsed = IsOverviewBoolTruthy(boldValue);
                if (boldUsed)
                {
                    labelRegex.Append("<b>");
                }

                if (typeValue != null)
                {
                    labelRegex.Append($"(?<{typeValue.Replace(" ", "")}>[A-Za-z0-9\\- \\']+)");
                }

                if (boldUsed)
                {
                    labelRegex.Append("</b>");
                }

                if (underlineUsed)
                {
                    labelRegex.Append("</u>");
                }

                if (itallicsUsed)
                {
                    labelRegex.Append("</i>");
                }

                var (_, postValue) = label.FirstOrDefault(x => x.Key == "post");
                if (postValue != null)
                {
                    labelRegex.Append(Regex.Escape(postValue));
                }

                if (colorValue != null)
                {
                    labelRegex.Append("</color>");
                }

                if (fontValue != null)
                {
                    labelRegex.Append("</font>");
                }

                var kvp = new KeyValuePair<string, string>(typeValue ?? "additionalText", labelRegex.ToString());
                labelRegexList.Add(kvp);
            }

            return labelRegexList;
        }
        catch (Exception ex)
        {
            throw new YamlException("Error parsing overview", ex);
        }
        
    }
    
    private static bool IsOverviewBoolTruthy(string? value)
    {
        // Value is not null, and value either equals "1", or "true"
        return value != null &&
               (value == "1" ||
                bool.TryParse(value, out var valueBool) && valueBool);
    }

    // Parse the overview yaml, extract the objects into key value pairs
    // Hard to keep a map in your mind of this, maybe its worth building a specific model?
    private static List<IEnumerable<KeyValuePair<string, string?>>> ExtractShipLabelSettings(
        IReadOnlyDictionary<string, object> parsedYaml)
    {
        var shipLabelOrderList = parsedYaml["shipLabelOrder"] as List<object> ??
                                 throw new Exception("shipLabelOrder is not in the overview file");

        var shipLabelOrder = shipLabelOrderList
            .Select(x => x as string ?? "additionalText")
            .ToList();

        var shipLabelList = parsedYaml["shipLabels"] as List<object> ??
                            throw new Exception("shipLabels is not in the overview file");

        var shipLabelListFancy = shipLabelList
            .Select(x =>
            {
                // Need to cast into an iterable list
                var itemList = x as List<object>
                               ?? throw new Exception("Unexpected value within the shipLabels");

                return itemList
                    // Each item in the list is a dictionary, cant cast as List<List<object>> for some reason
                    .Select(y => y as List<object>)!
                    // Compile our list of lists into key value pairs
                    .SelectMany<List<object>, KeyValuePair<string, string?>>(w =>
                    {
                        // The first key in each of these is not a list, so it casts to null, and we skip it
                        if (w == null) return new List<KeyValuePair<string, string?>>();

                        return w.Select(z =>
                        {
                            var zList = z as List<object>
                                        ?? throw new Exception("Unexpected value within the shipLabels");


                            var label = zList[0] as string
                                        ?? throw new Exception("Unexpected value within the shipLabels");
                            
                            // The color prop seems to use some RGB encoding or something
                            // We dont care about the value, but we need to know if there is one or not 
                            var value = (label == "color" && zList[1] != null) ? "YesAColorExists" : zList[1] as string;
                            return new KeyValuePair<string, string?>(label, value);
                        });
                    });
            });
            

        var resultList = new List<IEnumerable<KeyValuePair<string, string?>>>();
        foreach (var item in shipLabelOrder)
        {
            var shipLabel = shipLabelListFancy.FirstOrDefault(x => (x.FirstOrDefault(y => y.Key == "type").Value ?? "additionalText") == item);
            if (shipLabel == null)
            {
                Log.Warning("shipLabel was null. This isn't supposed to happen. It's possible parsing will have issues.");
                continue;
            }
            resultList.Add(shipLabel);
        }

        return resultList.Where(x =>
        {
            var typeValue = x.FirstOrDefault(y => y.Key == "type").Value;
            if (typeValue == "linebreak") return true;
            
            // Only include items that have `state == 1` which means they are visible
            var stateValue = x.FirstOrDefault(y => y.Key == "state").Value;
            if (!int.TryParse(stateValue, out var val)) return false;
            return val == 1;
        }).ToList();;
    }
}