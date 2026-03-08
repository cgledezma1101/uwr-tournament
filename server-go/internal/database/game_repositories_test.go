package database

import (
	"testing"

	"github.com/ing-bank/gormtestutil"
	. "github.com/onsi/gomega"
	"github.com/uwr-tournament/server-go/internal/models"
)

func TestGameRepository(t *testing.T) {
	db := gormtestutil.NewMemoryDatabase(t, gormtestutil.WithName("game_test"), gormtestutil.WithoutForeignKeys())
	db.AutoMigrate(&models.Club{}, &models.Team{}, &models.Tournament{}, &models.Stage{}, &models.Game{})

	clubRepo := NewClubRepository(db)
	teamRepo := NewTeamRepository(db)
	tournamentRepo := NewTournamentRepository(db)
	stageRepo := NewStageRepository(db)
	gameRepo := NewGameRepository(db)

	club := &models.Club{Name: "Game Club"}
	team := &models.Team{Name: "Game Team", ClubID: int64(club.ID)}
	tournament := &models.Tournament{Name: "Game Tournament"}
	stage := &models.Stage{Name: "Game Stage", TournamentID: tournament.ID}

	clubRepo.Create(club)
	teamRepo.Create(team)
	tournamentRepo.Create(tournament)
	stageRepo.Create(stage)

	t.Run("Create", func(t *testing.T) {
		g := NewWithT(t)
		game := &models.Game{
			BlueTeamID:   int(team.ID),
			WhiteTeamID:  int(team.ID),
			StageID:      stage.ID,
			Status:       "pending",
			WinningColor: "",
		}

		g.Expect(gameRepo.Create(game)).To(Succeed())
		g.Expect(game.ID).NotTo(BeZero())
	})

	t.Run("GetByID", func(t *testing.T) {
		g := NewWithT(t)
		game := &models.Game{
			BlueTeamID:   int(team.ID),
			WhiteTeamID:  int(team.ID),
			StageID:      stage.ID,
			Status:       "in_progress",
			WinningColor: "",
		}
		g.Expect(gameRepo.Create(game)).To(Succeed())

		retrieved, err := gameRepo.GetByID(game.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(retrieved).NotTo(BeNil())
		g.Expect(retrieved.Status).To(Equal("in_progress"))
	})

	t.Run("GetByID NotFound", func(t *testing.T) {
		g := NewWithT(t)
		retrieved, err := gameRepo.GetByID(99999)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(retrieved).To(BeNil())
	})

	t.Run("Update", func(t *testing.T) {
		g := NewWithT(t)
		game := &models.Game{
			BlueTeamID:   int(team.ID),
			WhiteTeamID:  int(team.ID),
			StageID:      stage.ID,
			Status:       "pending",
			WinningColor: "",
		}
		g.Expect(gameRepo.Create(game)).To(Succeed())

		game.Status = "completed"
		game.WinningColor = "blue"
		g.Expect(gameRepo.Update(game)).To(Succeed())

		retrieved, _ := gameRepo.GetByID(game.ID)
		g.Expect(retrieved.Status).To(Equal("completed"))
		g.Expect(retrieved.WinningColor).To(Equal("blue"))
	})

	t.Run("Delete", func(t *testing.T) {
		g := NewWithT(t)
		game := &models.Game{
			BlueTeamID:   int(team.ID),
			WhiteTeamID:  int(team.ID),
			StageID:      stage.ID,
			Status:       "pending",
			WinningColor: "",
		}
		g.Expect(gameRepo.Create(game)).To(Succeed())

		g.Expect(gameRepo.Delete(game.ID)).To(Succeed())

		retrieved, err := gameRepo.GetByID(game.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(retrieved).To(BeNil())
	})

	t.Run("List", func(t *testing.T) {
		g := NewWithT(t)
		games, err := gameRepo.List(0, 10)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(games)).To(BeNumerically(">=", 0))
	})

	t.Run("GetByStageID", func(t *testing.T) {
		g := NewWithT(t)
		games, err := gameRepo.GetByStageID(stage.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(games)).To(BeNumerically(">", 0))
	})
}

func TestGameEventRepository(t *testing.T) {
	db := gormtestutil.NewMemoryDatabase(t, gormtestutil.WithName("game_event_test"), gormtestutil.WithoutForeignKeys())
	db.AutoMigrate(&models.Club{}, &models.Team{}, &models.Tournament{}, &models.Stage{}, &models.Game{}, &models.GameEvent{})

	clubRepo := NewClubRepository(db)
	teamRepo := NewTeamRepository(db)
	tournamentRepo := NewTournamentRepository(db)
	stageRepo := NewStageRepository(db)
	gameRepo := NewGameRepository(db)
	eventRepo := NewGameEventRepository(db)

	club := &models.Club{Name: "Event Club"}
	team := &models.Team{Name: "Event Team", ClubID: int64(club.ID)}
	tournament := &models.Tournament{Name: "Event Tournament"}
	stage := &models.Stage{Name: "Event Stage", TournamentID: tournament.ID}
	game := &models.Game{
		BlueTeamID:   int(team.ID),
		WhiteTeamID:  int(team.ID),
		StageID:      stage.ID,
		Status:       "in_progress",
		WinningColor: "",
	}

	clubRepo.Create(club)
	teamRepo.Create(team)
	tournamentRepo.Create(tournament)
	stageRepo.Create(stage)
	gameRepo.Create(game)

	t.Run("Create", func(t *testing.T) {
		g := NewWithT(t)
		event := &models.GameEvent{
			Text:   "Goal scored",
			GameID: game.ID,
		}

		g.Expect(eventRepo.Create(event)).To(Succeed())
		g.Expect(event.ID).NotTo(BeZero())
	})

	t.Run("GetByGameID", func(t *testing.T) {
		g := NewWithT(t)
		event := &models.GameEvent{
			Text:   "Defense move",
			GameID: game.ID,
		}
		g.Expect(eventRepo.Create(event)).To(Succeed())

		events, err := eventRepo.GetByGameID(game.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(events)).To(BeNumerically(">", 0))
	})

	t.Run("Delete", func(t *testing.T) {
		g := NewWithT(t)
		event := &models.GameEvent{
			Text:   "Penalty",
			GameID: game.ID,
		}
		g.Expect(eventRepo.Create(event)).To(Succeed())

		g.Expect(eventRepo.Delete(event.ID)).To(Succeed())
	})
}

func TestScoreRepository(t *testing.T) {
	db := gormtestutil.NewMemoryDatabase(t, gormtestutil.WithName("score_test"), gormtestutil.WithoutForeignKeys())
	db.AutoMigrate(&models.User{}, &models.Club{}, &models.Team{}, &models.Player{}, &models.Tournament{}, &models.Stage{}, &models.Game{}, &models.Score{})

	userRepo := NewUserRepository(db)
	clubRepo := NewClubRepository(db)
	teamRepo := NewTeamRepository(db)
	playerRepo := NewPlayerRepository(db)
	tournamentRepo := NewTournamentRepository(db)
	stageRepo := NewStageRepository(db)
	gameRepo := NewGameRepository(db)
	scoreRepo := NewScoreRepository(db)

	user := &models.User{Email: "scorer@example.com", Name: "Scorer"}
	club := &models.Club{Name: "Score Club"}
	team := &models.Team{Name: "Score Team", ClubID: int64(club.ID)}
	tournament := &models.Tournament{Name: "Score Tournament"}
	stage := &models.Stage{Name: "Score Stage", TournamentID: tournament.ID}
	game := &models.Game{
		BlueTeamID:   int(team.ID),
		WhiteTeamID:  int(team.ID),
		StageID:      stage.ID,
		Status:       "completed",
		WinningColor: "blue",
	}

	userRepo.Create(user)
	clubRepo.Create(club)
	teamRepo.Create(team)
	tournamentRepo.Create(tournament)
	stageRepo.Create(stage)
	gameRepo.Create(game)

	isActive := true
	player := &models.Player{
		Number:   5,
		TeamID:   int(team.ID),
		UserID:   int64(user.ID),
		IsActive: &isActive,
	}
	playerRepo.Create(player)

	t.Run("Create", func(t *testing.T) {
		g := NewWithT(t)
		score := &models.Score{
			PlayerID: int(player.ID),
			GameID:   game.ID,
		}

		g.Expect(scoreRepo.Create(score)).To(Succeed())
		g.Expect(score.ID).NotTo(BeZero())
	})

	t.Run("GetByGameID", func(t *testing.T) {
		g := NewWithT(t)
		score := &models.Score{
			PlayerID: int(player.ID),
			GameID:   game.ID,
		}
		g.Expect(scoreRepo.Create(score)).To(Succeed())

		scores, err := scoreRepo.GetByGameID(game.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(scores)).To(BeNumerically(">", 0))
	})

	t.Run("GetByPlayerID", func(t *testing.T) {
		g := NewWithT(t)
		scores, err := scoreRepo.GetByPlayerID(int(player.ID))
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(scores)).To(BeNumerically(">", 0))
	})

	t.Run("Delete", func(t *testing.T) {
		g := NewWithT(t)
		score := &models.Score{
			PlayerID: int(player.ID),
			GameID:   game.ID,
		}
		g.Expect(scoreRepo.Create(score)).To(Succeed())

		g.Expect(scoreRepo.Delete(score.ID)).To(Succeed())
	})
}

func TestStageRepository(t *testing.T) {
	db := gormtestutil.NewMemoryDatabase(t, gormtestutil.WithName("stage_test"), gormtestutil.WithoutForeignKeys())
	db.AutoMigrate(&models.Tournament{}, &models.Stage{})

	tournamentRepo := NewTournamentRepository(db)
	stageRepo := NewStageRepository(db)

	tournament := &models.Tournament{Name: "Stage Tournament"}
	tournamentRepo.Create(tournament)

	t.Run("Create", func(t *testing.T) {
		g := NewWithT(t)
		stage := &models.Stage{
			Name:         "Group Stage",
			TournamentID: tournament.ID,
		}

		g.Expect(stageRepo.Create(stage)).To(Succeed())
		g.Expect(stage.ID).NotTo(BeZero())
	})

	t.Run("GetByID", func(t *testing.T) {
		g := NewWithT(t)
		stage := &models.Stage{
			Name:         "Knockout",
			TournamentID: tournament.ID,
		}
		g.Expect(stageRepo.Create(stage)).To(Succeed())

		retrieved, err := stageRepo.GetByID(stage.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(retrieved).NotTo(BeNil())
		g.Expect(retrieved.Name).To(Equal("Knockout"))
	})

	t.Run("Update", func(t *testing.T) {
		g := NewWithT(t)
		stage := &models.Stage{
			Name:         "Round of 16",
			TournamentID: tournament.ID,
		}
		g.Expect(stageRepo.Create(stage)).To(Succeed())

		stage.Name = "Round of 8"
		g.Expect(stageRepo.Update(stage)).To(Succeed())

		retrieved, _ := stageRepo.GetByID(stage.ID)
		g.Expect(retrieved.Name).To(Equal("Round of 8"))
	})

	t.Run("Delete", func(t *testing.T) {
		g := NewWithT(t)
		stage := &models.Stage{
			Name:         "Finals",
			TournamentID: tournament.ID,
		}
		g.Expect(stageRepo.Create(stage)).To(Succeed())

		g.Expect(stageRepo.Delete(stage.ID)).To(Succeed())

		retrieved, err := stageRepo.GetByID(stage.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(retrieved).To(BeNil())
	})

	t.Run("GetByTournamentID", func(t *testing.T) {
		g := NewWithT(t)
		stages, err := stageRepo.GetByTournamentID(tournament.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(stages)).To(BeNumerically(">", 0))
	})
}
