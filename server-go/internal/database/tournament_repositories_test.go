package database

import (
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"github.com/uwr-tournament/server-go/internal/models"
)

func TestTournamentRepository(t *testing.T) {
	tournamentRepo := NewTournamentRepository(Db)

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
	g := NewWithT(t)

	userRepo := NewUserRepository(Db)
	tournamentRepo := NewTournamentRepository(Db)
	adminRepo := NewTournamentAdminRepository(Db)

	user := &models.User{Email: "tadmin@example.com", Name: "Tournament Admin"}
	tournament := &models.Tournament{Name: "Admin Test Tournament"}

	g.Expect(userRepo.Create(user)).To(Succeed())
	g.Expect(tournamentRepo.Create(tournament)).To(Succeed())

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
	g := NewWithT(t)

	clubRepo := NewClubRepository(Db)
	tournamentRepo := NewTournamentRepository(Db)
	inviteRepo := NewTournamentInvitationRepository(Db)

	club := &models.Club{Name: "Invited Club"}
	tournament := &models.Tournament{Name: "Invite Tournament"}

	g.Expect(clubRepo.Create(club)).To(Succeed())
	g.Expect(tournamentRepo.Create(tournament)).To(Succeed())

	t.Run("Create, Get, Delete", func(t *testing.T) {
		g := NewWithT(t)
		invitation := &models.TournamentInvitation{
			ClubID:       club.ID,
			TournamentID: tournament.ID,
		}

		g.Expect(inviteRepo.Create(invitation)).To(Succeed())

		invitations, err := inviteRepo.GetByTournamentID(tournament.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(invitations)).To(BeNumerically(">", 0))

		invitations, err = inviteRepo.GetByClubID(club.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(invitations)).To(BeNumerically(">", 0))

		g.Expect(inviteRepo.Delete(invitation.ID)).To(Succeed())

		invitations, err = inviteRepo.GetByTournamentID(tournament.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(invitations)).To(BeZero())
	})
}

func TestTournamentTeamRepository(t *testing.T) {
	g := NewWithT(t)

	clubRepo := NewClubRepository(Db)
	teamRepo := NewTeamRepository(Db)
	tournamentRepo := NewTournamentRepository(Db)
	ttRepo := NewTournamentTeamRepository(Db)

	club := &models.Club{Name: "Team Club"}
	g.Expect(clubRepo.Create(club)).To(Succeed())
	team := &models.Team{Name: "Tournament Team", ClubID: int64(club.ID)}
	g.Expect(teamRepo.Create(team)).To(Succeed())
	tournament := &models.Tournament{Name: "Team Tournament"}
	g.Expect(tournamentRepo.Create(tournament)).To(Succeed())

	t.Run("Create, Get, Delete", func(t *testing.T) {
		g := NewWithT(t)
		tt := &models.TournamentTeam{
			TournamentID: tournament.ID,
			TeamID:       int(team.ID),
			Password:     "secret123",
		}
		g.Expect(ttRepo.Create(tt)).To(Succeed())

		tournamentTeams, err := ttRepo.GetByTournamentID(tournament.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(tournamentTeams)).To(BeNumerically(">", 0))

		tournamentTeams, err = ttRepo.GetByTeamID(int(team.ID))
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(tournamentTeams)).To(BeNumerically(">", 0))

		g.Expect(ttRepo.Delete(tt.ID)).To(Succeed())

		tournamentTeams, err = ttRepo.GetByTeamID(int(team.ID))
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(tournamentTeams)).To(BeZero())
	})
}
