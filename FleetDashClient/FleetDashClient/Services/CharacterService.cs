using FleetDashClient.Models;

namespace FleetDashClient.Services
{
    public class CharacterService : ICharacterService
    {
        public Task<Account[]> GetAccountList()
        {
            var accounts = new List<Account>();

            for (int i = 0; i < 10; i++)
            {
                AccountStatus status = AccountStatus.Ready;
                switch(i % 3)
                {
                    case 0:
                        status = AccountStatus.Ready; break;

                    case 1:
                        status = AccountStatus.ActivelyStreaming; break;

                    case 2:
                        status = AccountStatus.Error; break;

                }

                accounts.Add(new Account
                {
                    Status = status,
                    Name = "Hello",
                });
            }

            return Task.FromResult(accounts.ToArray());
        }
    }
}