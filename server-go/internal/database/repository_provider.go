package database

import (
	"github.com/uwr-tournament/server-go/internal/repositories"
	"gorm.io/gorm"
)

// RepositoryProvider manages all repository instances
type RepositoryProvider struct {
	Users              repositories.UserRepository
	Clubs              repositories.ClubRepository
	ClubAdmins         repositories.ClubAdminRepository
	ClubJoinRequests   repositories.ClubJoinRequestRepository
	Invitations        repositories.InvitationRepository
	UserClubs          repositories.UserClubRepository
	Teams              repositories.TeamRepository
	Players            repositories.PlayerRepository
	PlayerGames        repositories.PlayerGameRepository
	Games              repositories.GameRepository
	GameEvents         repositories.GameEventRepository
	Scores             repositories.ScoreRepository
	Stages             repositories.StageRepository
	Tournaments        repositories.TournamentRepository
	TournamentAdmins   repositories.TournamentAdminRepository
	TournamentInvitations repositories.TournamentInvitationRepository
	TournamentTeams    repositories.TournamentTeamRepository
}

// NewRepositoryProvider creates and initializes all repositories
func NewRepositoryProvider(db *gorm.DB) *RepositoryProvider {
	return &RepositoryProvider{
		Users:                  NewUserRepository(db),
		Clubs:                  NewClubRepository(db),
		ClubAdmins:             NewClubAdminRepository(db),
		ClubJoinRequests:       NewClubJoinRequestRepository(db),
		Invitations:            NewInvitationRepository(db),
		UserClubs:              NewUserClubRepository(db),
		Teams:                  NewTeamRepository(db),
		Players:                NewPlayerRepository(db),
		PlayerGames:            NewPlayerGameRepository(db),
		Games:                  NewGameRepository(db),
		GameEvents:             NewGameEventRepository(db),
		Scores:                 NewScoreRepository(db),
		Stages:                 NewStageRepository(db),
		Tournaments:            NewTournamentRepository(db),
		TournamentAdmins:       NewTournamentAdminRepository(db),
		TournamentInvitations: NewTournamentInvitationRepository(db),
		TournamentTeams:        NewTournamentTeamRepository(db),
	}
}
