using FleetDashClient.Models;

namespace FleetDashClient.Services;

public interface ICharacterService
{
    Task<Account[]> GetAccountList();
}