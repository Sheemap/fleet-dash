using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace FleetDashClient.Data.Migrations
{
    public partial class Configuration : Migration
    {
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.CreateTable(
                name: "Configurations",
                columns: table => new
                {
                    Id = table.Column<int>(type: "INTEGER", nullable: false)
                        .Annotation("Sqlite:Autoincrement", true),
                    LogDirectory = table.Column<string>(type: "TEXT", nullable: false),
                    Overview = table.Column<string>(type: "TEXT", nullable: true),
                    WindowHeight = table.Column<int>(type: "INTEGER", nullable: false),
                    WindowWidth = table.Column<int>(type: "INTEGER", nullable: false),
                    WindowPositionX = table.Column<int>(type: "INTEGER", nullable: false),
                    WindowPositionY = table.Column<int>(type: "INTEGER", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Configurations", x => x.Id);
                });
        }

        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "Configurations");
        }
    }
}
