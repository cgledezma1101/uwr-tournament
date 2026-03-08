package database

import (
	"testing"

	"github.com/ing-bank/gormtestutil"
	. "github.com/onsi/gomega"
	"github.com/uwr-tournament/server-go/internal/models"
)

func TestUserRepository(t *testing.T) {
	// Setup test database
	db := gormtestutil.NewMemoryDatabase(t, gormtestutil.WithName("user_test"), gormtestutil.WithoutForeignKeys())
	repo := NewUserRepository(db)

	t.Run("Create", func(t *testing.T) {
		g := NewWithT(t)
		user := &models.User{
			Email: "test@example.com",
			Name:  "Test User",
		}

		g.Expect(repo.Create(user)).To(Succeed())
		g.Expect(user.ID).NotTo(BeZero())
	})

	t.Run("GetByID", func(t *testing.T) {
		g := NewWithT(t)
		user := &models.User{
			Email: "getbyid@example.com",
			Name:  "GetByID User",
		}
		g.Expect(repo.Create(user)).To(Succeed())

		retrieved, err := repo.GetByID(user.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(retrieved).NotTo(BeNil())
		g.Expect(retrieved.Email).To(Equal("getbyid@example.com"))
	})

	t.Run("GetByID NotFound", func(t *testing.T) {
		g := NewWithT(t)
		retrieved, err := repo.GetByID(99999)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(retrieved).To(BeNil())
	})

	t.Run("GetByEmail", func(t *testing.T) {
		g := NewWithT(t)
		user := &models.User{
			Email: "unique@example.com",
			Name:  "Email User",
		}
		g.Expect(repo.Create(user)).To(Succeed())

		retrieved, err := repo.GetByEmail("unique@example.com")
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(retrieved).NotTo(BeNil())
		g.Expect(retrieved.Name).To(Equal("Email User"))
	})

	t.Run("GetByEmail NotFound", func(t *testing.T) {
		g := NewWithT(t)
		retrieved, err := repo.GetByEmail("nonexistent@example.com")
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(retrieved).To(BeNil())
	})

	t.Run("Update", func(t *testing.T) {
		g := NewWithT(t)
		user := &models.User{
			Email: "update@example.com",
			Name:  "Original Name",
		}
		g.Expect(repo.Create(user)).To(Succeed())

		user.Name = "Updated Name"
		g.Expect(repo.Update(user)).To(Succeed())

		retrieved, _ := repo.GetByID(user.ID)
		g.Expect(retrieved.Name).To(Equal("Updated Name"))
	})

	t.Run("Delete", func(t *testing.T) {
		g := NewWithT(t)
		user := &models.User{
			Email: "delete@example.com",
			Name:  "To Delete",
		}
		g.Expect(repo.Create(user)).To(Succeed())

		g.Expect(repo.Delete(user.ID)).To(Succeed())

		retrieved, err := repo.GetByID(user.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(retrieved).To(BeNil())
	})

	t.Run("List", func(t *testing.T) {
		g := NewWithT(t)
		for i := 1; i <= 5; i++ {
			g.Expect(repo.Create(&models.User{
				Email: "list" + string(rune(i)) + "@example.com",
				Name:  "List User " + string(rune(i)),
			})).To(Succeed())
		}

		users, err := repo.List(0, 3)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(users)).To(BeNumerically("<=", 3))
	})

	t.Run("List with pagination", func(t *testing.T) {
		g := NewWithT(t)
		users1, err := repo.List(0, 2)
		g.Expect(err).NotTo(HaveOccurred())

		users2, err := repo.List(2, 2)
		g.Expect(err).NotTo(HaveOccurred())

		// Ensure we got different results
		if len(users1) > 0 && len(users2) > 0 {
			g.Expect(users1[0].ID).NotTo(Equal(users2[0].ID))
		}
	})
}

func TestUserClubRepository(t *testing.T) {
	db := gormtestutil.NewMemoryDatabase(t, gormtestutil.WithName("user_club_test"), gormtestutil.WithoutForeignKeys())
	db.AutoMigrate(&models.User{}, &models.Club{}, &models.UserClub{})

	userRepo := NewUserRepository(db)
	clubRepo := NewClubRepository(db)
	userClubRepo := NewUserClubRepository(db)

	user := &models.User{Email: "member@example.com", Name: "Member"}
	club := &models.Club{Name: "Member Club"}
	g := NewWithT(t)
	g.Expect(userRepo.Create(user)).To(Succeed())
	g.Expect(clubRepo.Create(club)).To(Succeed())

	t.Run("Create", func(t *testing.T) {
		g := NewWithT(t)
		userClub := &models.UserClub{
			UserID: user.ID,
			ClubID: club.ID,
		}

		g.Expect(userClubRepo.Create(userClub)).To(Succeed())
	})

	t.Run("GetByUserID", func(t *testing.T) {
		g := NewWithT(t)
		userClubs, err := userClubRepo.GetByUserID(user.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(userClubs)).To(BeNumerically(">", 0))
	})

	t.Run("GetByClubID", func(t *testing.T) {
		g := NewWithT(t)
		userClubs, err := userClubRepo.GetByClubID(club.ID)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(len(userClubs)).To(BeNumerically(">", 0))
	})

	t.Run("DeleteByUserAndClub", func(t *testing.T) {
		g := NewWithT(t)
		g.Expect(userClubRepo.DeleteByUserAndClub(user.ID, club.ID)).To(Succeed())

		userClubs, _ := userClubRepo.GetByUserID(user.ID)
		g.Expect(len(userClubs)).To(Equal(0))
	})
}
