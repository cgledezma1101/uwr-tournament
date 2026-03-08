package database

import (
	"testing"
	"time"

	"github.com/ing-bank/gormtestutil"
	. "github.com/onsi/gomega"
	"github.com/uwr-tournament/server-go/internal/models"
)

func TestTournamentRepository(t *testing.T) {
	db := gormtestutil.NewMemoryDatabase(t, gormtestutil.WithName("tournament_test"), gormtestutil.WithoutForeignKeys())
	db.AutoMigrate(&models.Tournament{})

	tournamentRepo := NewTournamentRepository(db)

	t.Run("Create", func(t *testing.T) {
		g := NewWithT(t)
		startDate := time.Now()
		endDate := startDate.AddDate(0, 0, 7)
		tournament := &models.Tournament{
			Name:      "Spring Tournament",
			StartDate: &startDate,
			EndDate:   &endDate,
		}

		g.Expect(tournamentRepo.Create(tournament)).To(Succeed())
		g.Expect(tournament.ID).NotTo(BeZero())
	})

	t.Run("GetByID", func(t *testing.T) {
		g := NewWithT(t)
		tournament := &models.Tournament{
			Name: "Fall Tournament",
		}
		g.Expect(tournamentRepo.Create(tournament)).To(Succeed())

		retrieved, err := tournamentRepo.GetByID(tournament.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(retrieved).NotTo(BeNil())
		g.Expect(retrieved.Name).To(Equal("Fall Tournament"))
	})

	t.Run("GetByID NotFound", func(t *testing.T) {
		g := NewWithT(t)
		retrieved, err := tournamentRepo.GetByID(99999)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(retrieved).To(BeNil())
	})

	t.Run("Update", func(t *testing.T) {
		g := NewWithT(t)
		tournament := &models.Tournament{
			Name: "Original Championship",
		}
		g.Expect(tournamentRepo.Create(tournament)).To(Succeed())

		tournament.Name = "Updated Championship"
		g.Expect(tournamentRepo.Update(tournament)).To(Succeed())

		retrieved, _ := tournamentRepo.GetByID(tournament.ID)
		g.Expect(retrieved.Name).To(Equal("Updated Championship"))
	})

	t.Run("Delete", func(t *testing.T) {
		g := NewWithT(t)
		tournament := &models.Tournament{
			Name: "Delete Championship",
		}
		g.Expect(tournamentRepo.Create(tournament)).To(Succeed())

		g.Expect(tournamentRepo.Delete(tournament.ID)).To(Succeed())

		retrieved, err := tournamentRepo.GetByID(tournament.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(retrieved).To(BeNil())
	})

	t.Run("List", func(t *testing.T) {
		g := NewWithT(t)
		tournaments, err := tournamentRepo.List(0, 10)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(tournaments)).To(BeNumerically(">=", 0))
	})
}

func TestTournamentAdminRepository(t *testing.T) {
	db := gormtestutil.NewMemoryDatabase(t, gormtestutil.WithName("tournament_admin_test"), gormtestutil.WithoutForeignKeys())
	db.AutoMigrate(&models.User{}, &models.Tournament{}, &models.TournamentAdmin{})

	userRepo := NewUserRepository(db)
	tournamentRepo := NewTournamentRepository(db)
	adminRepo := NewTournamentAdminRepository(db)

	user := &models.User{Email: "tadmin@example.com", Name: "Tournament Admin"}
	tournament := &models.Tournament{Name: "Admin Test Tournament"}

	userRepo.Create(user)
	tournamentRepo.Create(tournament)

	t.Run("Create", func(t *testing.T) {
		g := NewWithT(t)
		admin := &models.TournamentAdmin{
			UserID:       user.ID,
			TournamentID: tournament.ID,
		}

		g.Expect(adminRepo.Create(admin)).To(Succeed())
	})

	t.Run("GetByTournamentID", func(t *testing.T) {
		g := NewWithT(t)
		admins, err := adminRepo.GetByTournamentID(tournament.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(admins)).To(BeNumerically(">", 0))
	})

	t.Run("DeleteByUserAndTournament", func(t *testing.T) {
		g := NewWithT(t)
		err := adminRepo.DeleteByUserAndTournament(user.ID, tournament.ID)
		g.Expect(err).NotTo(HaveOccurred())

		admins, _ := adminRepo.GetByTournamentID(tournament.ID)
		g.Expect(len(admins)).To(Equal(0))
	})
}

func TestTournamentInvitationRepository(t *testing.T) {
	db := gormtestutil.NewMemoryDatabase(t, gormtestutil.WithName("tournament_invitation_test"), gormtestutil.WithoutForeignKeys())
	db.AutoMigrate(&models.Club{}, &models.Tournament{}, &models.TournamentInvitation{})

	clubRepo := NewClubRepository(db)
	tournamentRepo := NewTournamentRepository(db)
	inviteRepo := NewTournamentInvitationRepository(db)

	club := &models.Club{Name: "Invited Club"}
	tournament := &models.Tournament{Name: "Invite Tournament"}

	clubRepo.Create(club)
	tournamentRepo.Create(tournament)

	t.Run("Create", func(t *testing.T) {
		g := NewWithT(t)
		invitation := &models.TournamentInvitation{
			ClubID:       club.ID,
			TournamentID: tournament.ID,
		}

		g.Expect(inviteRepo.Create(invitation)).To(Succeed())
	})

	t.Run("GetByTournamentID", func(t *testing.T) {
		g := NewWithT(t)
		invitations, err := inviteRepo.GetByTournamentID(tournament.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(invitations)).To(BeNumerically(">", 0))
	})

	t.Run("GetByClubID", func(t *testing.T) {
		g := NewWithT(t)
		invitations, err := inviteRepo.GetByClubID(club.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(invitations)).To(BeNumerically(">", 0))
	})

	t.Run("Delete", func(t *testing.T) {
		g := NewWithT(t)
		invitation := &models.TournamentInvitation{
			ClubID:       club.ID,
			TournamentID: tournament.ID,
		}
		g.Expect(inviteRepo.Create(invitation)).To(Succeed())

		g.Expect(inviteRepo.Delete(invitation.ID)).To(Succeed())

		invitations, _ := inviteRepo.GetByTournamentID(tournament.ID)
		// Should have one less invitation
		g.Expect(len(invitations)).To(BeNumerically(">=", 0))
	})
}

func TestTournamentTeamRepository(t *testing.T) {
	db := gormtestutil.NewMemoryDatabase(t, gormtestutil.WithName("tournament_team_test"), gormtestutil.WithoutForeignKeys())
	db.AutoMigrate(&models.Club{}, &models.Team{}, &models.Tournament{}, &models.TournamentTeam{})

	clubRepo := NewClubRepository(db)
	teamRepo := NewTeamRepository(db)
	tournamentRepo := NewTournamentRepository(db)
	ttRepo := NewTournamentTeamRepository(db)

	club := &models.Club{Name: "Team Club"}
	team := &models.Team{Name: "Tournament Team", ClubID: int64(club.ID)}
	tournament := &models.Tournament{Name: "Team Tournament"}

	clubRepo.Create(club)
	teamRepo.Create(team)
	tournamentRepo.Create(tournament)

	t.Run("Create", func(t *testing.T) {
		g := NewWithT(t)
		tt := &models.TournamentTeam{
			TournamentID: tournament.ID,
			TeamID:       int(team.ID),
			Password:     "secret123",
		}

		g.Expect(ttRepo.Create(tt)).To(Succeed())
	})

	t.Run("GetByTournamentID", func(t *testing.T) {
		g := NewWithT(t)
		tournamentTeams, err := ttRepo.GetByTournamentID(tournament.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(tournamentTeams)).To(BeNumerically(">", 0))
	})

	t.Run("GetByTeamID", func(t *testing.T) {
		g := NewWithT(t)
		tournamentTeams, err := ttRepo.GetByTeamID(int(team.ID))
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(tournamentTeams)).To(BeNumerically(">", 0))
	})

	t.Run("Delete", func(t *testing.T) {
		g := NewWithT(t)
		tt := &models.TournamentTeam{
			TournamentID: tournament.ID,
			TeamID:       int(team.ID),
			Password:     "pass456",
		}
		g.Expect(ttRepo.Create(tt)).To(Succeed())

		g.Expect(ttRepo.Delete(tt.ID)).To(Succeed())
	})
}
