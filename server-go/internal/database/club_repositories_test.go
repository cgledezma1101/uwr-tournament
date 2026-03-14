package database

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/uwr-tournament/server-go/internal/models"
)

func TestClubRepository(t *testing.T) {
	repo := NewClubRepository(Db)

	t.Run("Create", func(t *testing.T) {
		g := NewWithT(t)
		club := &models.Club{
			Name: "Test Club",
		}

		g.Expect(repo.Create(club)).To(Succeed())
		g.Expect(club.ID).NotTo(BeZero())
	})

	t.Run("GetByID", func(t *testing.T) {
		g := NewWithT(t)
		club := &models.Club{
			Name: "Get Club",
		}
		g.Expect(repo.Create(club)).To(Succeed())

		retrieved, err := repo.GetByID(club.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(retrieved).NotTo(BeNil())
		g.Expect(retrieved.Name).To(Equal("Get Club"))
	})

	t.Run("GetByID NotFound", func(t *testing.T) {
		g := NewWithT(t)
		retrieved, err := repo.GetByID(99999)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(retrieved).To(BeNil())
	})

	t.Run("Update", func(t *testing.T) {
		g := NewWithT(t)
		club := &models.Club{
			Name: "Original Club",
		}
		g.Expect(repo.Create(club)).To(Succeed())

		club.Name = "Updated Club"
		g.Expect(repo.Update(club)).To(Succeed())

		retrieved, _ := repo.GetByID(club.ID)
		g.Expect(retrieved.Name).To(Equal("Updated Club"))
	})

	t.Run("Delete", func(t *testing.T) {
		g := NewWithT(t)
		club := &models.Club{
			Name: "Delete Club",
		}
		g.Expect(repo.Create(club)).To(Succeed())

		g.Expect(repo.Delete(club.ID)).To(Succeed())

		retrieved, err := repo.GetByID(club.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(retrieved).To(BeNil())
	})

	t.Run("List", func(t *testing.T) {
		g := NewWithT(t)
		for i := 1; i <= 5; i++ {
			club := &models.Club{
				Name: "Club " + string(rune(i+48)),
			}
			g.Expect(repo.Create(club)).To(Succeed())
		}

		clubs, err := repo.List(0, 10)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(clubs)).To(BeNumerically(">", 0))
	})
}

func TestClubAdminRepository(t *testing.T) {
	g := NewWithT(t)

	userRepo := NewUserRepository(Db)
	clubRepo := NewClubRepository(Db)
	adminRepo := NewClubAdminRepository(Db)

	user := &models.User{Email: "admin@example.com", Name: "Admin"}
	club := &models.Club{Name: "Test Club"}
	g.Expect(userRepo.Create(user)).To(Succeed())
	g.Expect(clubRepo.Create(club)).To(Succeed())

	t.Run("Create", func(t *testing.T) {
		g := NewWithT(t)
		admin := &models.ClubAdmin{
			UserID: user.ID,
			ClubID: club.ID,
		}

		g.Expect(adminRepo.Create(admin)).To(Succeed())
	})

	t.Run("GetByClubID", func(t *testing.T) {
		g := NewWithT(t)
		admins, err := adminRepo.GetByClubID(club.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(admins)).To(BeNumerically(">", 0))
	})

	t.Run("DeleteByUserAndClub", func(t *testing.T) {
		g := NewWithT(t)
		g.Expect(adminRepo.DeleteByUserAndClub(user.ID, club.ID)).To(Succeed())

		admins, _ := adminRepo.GetByClubID(club.ID)
		g.Expect(len(admins)).To(Equal(0))
	})
}

func TestClubJoinRequestRepository(t *testing.T) {
	g := NewWithT(t)

	userRepo := NewUserRepository(Db)
	clubRepo := NewClubRepository(Db)
	requestRepo := NewClubJoinRequestRepository(Db)

	user := &models.User{Email: "requester@example.com", Name: "Requester"}
	club := &models.Club{Name: "Target Club"}
	g.Expect(userRepo.Create(user)).To(Succeed())
	g.Expect(clubRepo.Create(club)).To(Succeed())

	t.Run("Create, Get, Delete", func(t *testing.T) {
		g := NewWithT(t)
		request := &models.ClubJoinRequest{
			UserID: user.ID,
			ClubID: club.ID,
		}

		g.Expect(requestRepo.Create(request)).To(Succeed())

		retrieved, err := requestRepo.GetByID(request.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(retrieved).NotTo(BeNil())
		g.Expect(retrieved.UserID).To(Equal(user.ID))

		requests, err := requestRepo.GetByClubID(club.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(requests)).To(BeNumerically(">", 0))

		g.Expect(requestRepo.Delete(request.ID)).To(Succeed())

		retrieved, err = requestRepo.GetByID(request.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(retrieved).To(BeNil())
	})
}

func TestInvitationRepository(t *testing.T) {
	g := NewWithT(t)

	userRepo := NewUserRepository(Db)
	clubRepo := NewClubRepository(Db)
	inviteRepo := NewInvitationRepository(Db)

	user := &models.User{Email: "invited@example.com", Name: "Invited"}
	club := &models.Club{Name: "Host Club"}
	g.Expect(userRepo.Create(user)).To(Succeed())
	g.Expect(clubRepo.Create(club)).To(Succeed())

	t.Run("Create, Get, Delete", func(t *testing.T) {
		g := NewWithT(t)
		isAdmin := true
		invitation := &models.Invitation{
			UserID:  user.ID,
			ClubID:  club.ID,
			IsAdmin: &isAdmin,
		}

		g.Expect(inviteRepo.Create(invitation)).To(Succeed())

		retrieved, err := inviteRepo.GetByID(invitation.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(retrieved).NotTo(BeNil())
		g.Expect(*retrieved.IsAdmin).To(Equal(true))

		invitations, err := inviteRepo.GetByUserID(user.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(invitations)).To(BeNumerically(">", 0))

		invitations, err = inviteRepo.GetByClubID(club.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(invitations)).To(BeNumerically(">", 0))

		g.Expect(inviteRepo.Delete(invitation.ID)).To(Succeed())

		retrieved, err = inviteRepo.GetByID(invitation.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(retrieved).To(BeNil())
	})
}
