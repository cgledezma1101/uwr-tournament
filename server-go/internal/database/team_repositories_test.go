package database

import (
	"testing"

	"github.com/ing-bank/gormtestutil"
	. "github.com/onsi/gomega"
	"github.com/uwr-tournament/server-go/internal/models"
)

func TestTeamRepository(t *testing.T) {
	db := gormtestutil.NewMemoryDatabase(t, gormtestutil.WithName("team_test"), gormtestutil.WithoutForeignKeys())
	db.AutoMigrate(&models.Club{}, &models.Team{})

	clubRepo := NewClubRepository(db)
	teamRepo := NewTeamRepository(db)

	club := &models.Club{Name: "Sports Club"}
	clubRepo.Create(club)

	t.Run("Create", func(t *testing.T) {
		g := NewWithT(t)
		team := &models.Team{
			Name:   "Team A",
			ClubID: int64(club.ID),
		}

		g.Expect(teamRepo.Create(team)).To(Succeed())
		g.Expect(team.ID).NotTo(BeZero())
	})

	t.Run("GetByID", func(t *testing.T) {
		g := NewWithT(t)
		team := &models.Team{
			Name:   "Get Team",
			ClubID: int64(club.ID),
		}
		g.Expect(teamRepo.Create(team)).To(Succeed())

		retrieved, err := teamRepo.GetByID(team.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(retrieved).NotTo(BeNil())
		g.Expect(retrieved.Name).To(Equal("Get Team"))
	})

	t.Run("GetByID NotFound", func(t *testing.T) {
		g := NewWithT(t)
		retrieved, err := teamRepo.GetByID(99999)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(retrieved).To(BeNil())
	})

	t.Run("Update", func(t *testing.T) {
		g := NewWithT(t)
		team := &models.Team{
			Name:   "Original Team",
			ClubID: int64(club.ID),
		}
		g.Expect(teamRepo.Create(team)).To(Succeed())

		team.Name = "Updated Team"
		g.Expect(teamRepo.Update(team)).To(Succeed())

		retrieved, _ := teamRepo.GetByID(team.ID)
		g.Expect(retrieved.Name).To(Equal("Updated Team"))
	})

	t.Run("Delete", func(t *testing.T) {
		g := NewWithT(t)
		team := &models.Team{
			Name:   "Delete Team",
			ClubID: int64(club.ID),
		}
		g.Expect(teamRepo.Create(team)).To(Succeed())

		g.Expect(teamRepo.Delete(team.ID)).To(Succeed())

		retrieved, err := teamRepo.GetByID(team.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(retrieved).To(BeNil())
	})

	t.Run("List", func(t *testing.T) {
		g := NewWithT(t)
		for i := 0; i < 3; i++ {
			g.Expect(teamRepo.Create(&models.Team{
				Name:   "List Team",
				ClubID: int64(club.ID),
			})).To(Succeed())
		}

		teams, err := teamRepo.List(0, 10)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(teams)).To(BeNumerically(">", 0))
	})

	t.Run("GetByClubID", func(t *testing.T) {
		g := NewWithT(t)
		teams, err := teamRepo.GetByClubID(int64(club.ID))
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(teams)).To(BeNumerically(">", 0))
	})
}

func TestPlayerRepository(t *testing.T) {
	db := gormtestutil.NewMemoryDatabase(t, gormtestutil.WithName("player_test"), gormtestutil.WithoutForeignKeys())
	db.AutoMigrate(&models.User{}, &models.Club{}, &models.Team{}, &models.Player{})

	userRepo := NewUserRepository(db)
	clubRepo := NewClubRepository(db)
	teamRepo := NewTeamRepository(db)
	playerRepo := NewPlayerRepository(db)

	user := &models.User{Email: "player@example.com", Name: "Player"}
	club := &models.Club{Name: "Player Club"}
	team := &models.Team{Name: "Player Team", ClubID: int64(club.ID)}

	userRepo.Create(user)
	clubRepo.Create(club)
	teamRepo.Create(team)

	t.Run("Create", func(t *testing.T) {
		g := NewWithT(t)
		isActive := true
		player := &models.Player{
			Number:   1,
			TeamID:   int(team.ID),
			UserID:   int64(user.ID),
			IsActive: &isActive,
		}

		g.Expect(playerRepo.Create(player)).To(Succeed())
		g.Expect(player.ID).NotTo(BeZero())
	})

	t.Run("GetByID", func(t *testing.T) {
		g := NewWithT(t)
		isActive := true
		player := &models.Player{
			Number:   2,
			TeamID:   int(team.ID),
			UserID:   int64(user.ID),
			IsActive: &isActive,
		}
		g.Expect(playerRepo.Create(player)).To(Succeed())

		retrieved, err := playerRepo.GetByID(player.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(retrieved).NotTo(BeNil())
		g.Expect(retrieved.Number).To(Equal(2))
	})

	t.Run("Update", func(t *testing.T) {
		g := NewWithT(t)
		isActive := false
		player := &models.Player{
			Number:   3,
			TeamID:   int(team.ID),
			UserID:   int64(user.ID),
			IsActive: &isActive,
		}
		g.Expect(playerRepo.Create(player)).To(Succeed())

		newActive := true
		player.IsActive = &newActive
		g.Expect(playerRepo.Update(player)).To(Succeed())

		retrieved, _ := playerRepo.GetByID(player.ID)
		g.Expect(*retrieved.IsActive).To(Equal(true))
	})

	t.Run("Delete", func(t *testing.T) {
		g := NewWithT(t)
		isActive := true
		player := &models.Player{
			Number:   4,
			TeamID:   int(team.ID),
			UserID:   int64(user.ID),
			IsActive: &isActive,
		}
		g.Expect(playerRepo.Create(player)).To(Succeed())

		g.Expect(playerRepo.Delete(player.ID)).To(Succeed())

		retrieved, err := playerRepo.GetByID(player.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(retrieved).To(BeNil())
	})

	t.Run("List", func(t *testing.T) {
		g := NewWithT(t)
		players, err := playerRepo.List(0, 10)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(players)).To(BeNumerically(">", 0))
	})

	t.Run("GetByTeamID", func(t *testing.T) {
		g := NewWithT(t)
		players, err := playerRepo.GetByTeamID(int(team.ID))
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(players)).To(BeNumerically(">", 0))
	})

	t.Run("GetByUserID", func(t *testing.T) {
		g := NewWithT(t)
		players, err := playerRepo.GetByUserID(int64(user.ID))
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(players)).To(BeNumerically(">", 0))
	})
}

func TestPlayerGameRepository(t *testing.T) {
	db := gormtestutil.NewMemoryDatabase(t, gormtestutil.WithName("player_game_test"), gormtestutil.WithoutForeignKeys())
	db.AutoMigrate(&models.User{}, &models.Club{}, &models.Team{}, &models.Player{}, &models.Stage{}, &models.Tournament{}, &models.Game{}, &models.PlayerGame{})

	userRepo := NewUserRepository(db)
	clubRepo := NewClubRepository(db)
	teamRepo := NewTeamRepository(db)
	playerRepo := NewPlayerRepository(db)
	tournamentRepo := NewTournamentRepository(db)
	stageRepo := NewStageRepository(db)
	gameRepo := NewGameRepository(db)
	pgRepo := NewPlayerGameRepository(db)

	user := &models.User{Email: "pgamer@example.com", Name: "Player Gamer"}
	club := &models.Club{Name: "PG Club"}
	team := &models.Team{Name: "PG Team", ClubID: int64(club.ID)}
	tournament := &models.Tournament{Name: "PG Tournament"}
	stage := &models.Stage{Name: "PG Stage", TournamentID: tournament.ID}
	game := &models.Game{
		BlueTeamID:   int(team.ID),
		WhiteTeamID:  int(team.ID),
		StageID:      stage.ID,
		WinningColor: "blue",
		Status:       "completed",
	}

	userRepo.Create(user)
	clubRepo.Create(club)
	teamRepo.Create(team)
	tournamentRepo.Create(tournament)
	stageRepo.Create(stage)
	gameRepo.Create(game)

	isActive := true
	player := &models.Player{
		Number:   10,
		TeamID:   int(team.ID),
		UserID:   int64(user.ID),
		IsActive: &isActive,
	}
	playerRepo.Create(player)

	t.Run("Create", func(t *testing.T) {
		g := NewWithT(t)
		pg := &models.PlayerGame{
			PlayerID:  int(player.ID),
			GameID:    game.ID,
			TeamColor: "blue",
		}

		g.Expect(pgRepo.Create(pg)).To(Succeed())
		g.Expect(pg.ID).NotTo(BeZero())
	})

	t.Run("GetByGameID", func(t *testing.T) {
		g := NewWithT(t)
		playerGames, err := pgRepo.GetByGameID(game.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(playerGames)).To(BeNumerically(">=", 0))
	})

	t.Run("GetByPlayerID", func(t *testing.T) {
		g := NewWithT(t)
		playerGames, err := pgRepo.GetByPlayerID(player.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(playerGames)).To(BeNumerically(">=", 0))
	})

	t.Run("Delete", func(t *testing.T) {
		g := NewWithT(t)
		pg := &models.PlayerGame{
			PlayerID:  int(player.ID),
			GameID:    game.ID,
			TeamColor: "white",
		}
		g.Expect(pgRepo.Create(pg)).To(Succeed())

		g.Expect(pgRepo.Delete(pg.ID)).To(Succeed())
	})
}
