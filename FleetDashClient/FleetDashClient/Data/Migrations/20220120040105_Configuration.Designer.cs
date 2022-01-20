﻿// <auto-generated />
using System;
using FleetDashClient.Data;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Infrastructure;
using Microsoft.EntityFrameworkCore.Migrations;
using Microsoft.EntityFrameworkCore.Storage.ValueConversion;

#nullable disable

namespace FleetDashClient.Data.Migrations
{
    [DbContext(typeof(DataContext))]
    [Migration("20220120040105_Configuration")]
    partial class Configuration
    {
        protected override void BuildTargetModel(ModelBuilder modelBuilder)
        {
#pragma warning disable 612, 618
            modelBuilder.HasAnnotation("ProductVersion", "6.0.1");

            modelBuilder.Entity("FleetDashClient.Models.Character", b =>
                {
                    b.Property<string>("Id")
                        .HasColumnType("TEXT");

                    b.Property<string>("Name")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.HasKey("Id");

                    b.ToTable("Characters");
                });

            modelBuilder.Entity("FleetDashClient.Models.Configuration", b =>
                {
                    b.Property<int>("Id")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("INTEGER");

                    b.Property<string>("LogDirectory")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.Property<string>("Overview")
                        .HasColumnType("TEXT");

                    b.Property<int>("WindowHeight")
                        .HasColumnType("INTEGER");

                    b.Property<int>("WindowPositionX")
                        .HasColumnType("INTEGER");

                    b.Property<int>("WindowPositionY")
                        .HasColumnType("INTEGER");

                    b.Property<int>("WindowWidth")
                        .HasColumnType("INTEGER");

                    b.HasKey("Id");

                    b.ToTable("Configurations");
                });

            modelBuilder.Entity("FleetDashClient.Models.Token", b =>
                {
                    b.Property<string>("CharacterId")
                        .HasColumnType("TEXT");

                    b.Property<string>("AccessToken")
                        .HasColumnType("TEXT");

                    b.Property<DateTimeOffset>("ExpiresAt")
                        .HasColumnType("TEXT");

                    b.Property<int>("ExpiresIn")
                        .HasColumnType("INTEGER");

                    b.Property<string>("RefreshToken")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.Property<string>("TokenType")
                        .IsRequired()
                        .HasColumnType("TEXT");

                    b.HasKey("CharacterId", "AccessToken");

                    b.HasIndex("ExpiresAt");

                    b.ToTable("Tokens");
                });

            modelBuilder.Entity("FleetDashClient.Models.Token", b =>
                {
                    b.HasOne("FleetDashClient.Models.Character", "Character")
                        .WithMany("Tokens")
                        .HasForeignKey("CharacterId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();

                    b.Navigation("Character");
                });

            modelBuilder.Entity("FleetDashClient.Models.Character", b =>
                {
                    b.Navigation("Tokens");
                });
#pragma warning restore 612, 618
        }
    }
}
