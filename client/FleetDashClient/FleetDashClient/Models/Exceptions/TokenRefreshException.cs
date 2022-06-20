namespace FleetDashClient.Models.Exceptions;

public class TokenRefreshException : Exception
{
    public TokenRefreshException(string message) : base(message)
    {
    }
}