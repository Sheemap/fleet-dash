using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace FleetDashClient.Models.Events
{
    public class LogFileReadEventArgs : EventArgs
    {
        public LogFileReadEventArgs(string characterID, byte[] content)
        {
            CharacterID = characterID;
            Content = content;
        }

        public string CharacterID { get; set; }
        public byte[] Content { get; set; }
    }
}
