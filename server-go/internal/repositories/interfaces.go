package repositories

//go:generate mockgen -source=interfaces.go -destination=mocks/interfaces.go -package=mocks
import "github.com/uwr-tournament/server-go/internal/models"

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id int) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(id int) error
	List(offset, limit int) ([]models.User, error)
}

type ClubRepository interface {
	Create(club *models.Club) error
	GetByID(id int) (*models.Club, error)
	Update(club *models.Club) error
	Delete(id int) error
	List(offset, limit int) ([]models.Club, error)
}

type ClubAdminRepository interface {
	Create(admin *models.ClubAdmin) error
	GetByClubID(clubID int) ([]models.ClubAdmin, error)
	DeleteByUserAndClub(userID, clubID int) error
}

type ClubJoinRequestRepository interface {
	Create(request *models.ClubJoinRequest) error
	GetByID(id int) (*models.ClubJoinRequest, error)
	GetByClubID(clubID int) ([]models.ClubJoinRequest, error)
	Delete(id int) error
}

type InvitationRepository interface {
	Create(invitation *models.Invitation) error
	GetByID(id int) (*models.Invitation, error)
	GetByUserID(userID int) ([]models.Invitation, error)
	GetByClubID(clubID int) ([]models.Invitation, error)
	Delete(id int) error
}

type UserClubRepository interface {
	Create(userClub *models.UserClub) error
	GetByUserID(userID int) ([]models.UserClub, error)
	GetByClubID(clubID int) ([]models.UserClub, error)
	DeleteByUserAndClub(userID, clubID int) error
}

type TeamRepository interface {
	Create(team *models.Team) error
	GetByID(id int64) (*models.Team, error)
	Update(team *models.Team) error
	Delete(id int64) error
	List(offset, limit int) ([]models.Team, error)
	GetByClubID(clubID int64) ([]models.Team, error)
}

type PlayerRepository interface {
	Create(player *models.Player) error
	GetByID(id int64) (*models.Player, error)
	Update(player *models.Player) error
	Delete(id int64) error
	List(offset, limit int) ([]models.Player, error)
	GetByTeamID(teamID int) ([]models.Player, error)
	GetByUserID(userID int64) ([]models.Player, error)
}

type PlayerGameRepository interface {
	Create(playerGame *models.PlayerGame) error
	GetByGameID(gameID int) ([]models.PlayerGame, error)
	GetByPlayerID(playerID int64) ([]models.PlayerGame, error)
	Delete(id int64) error
}

type GameRepository interface {
	Create(game *models.Game) error
	GetByID(id int) (*models.Game, error)
	Update(game *models.Game) error
	Delete(id int) error
	List(offset, limit int) ([]models.Game, error)
	GetByStageID(stageID int) ([]models.Game, error)
}

type GameEventRepository interface {
	Create(event *models.GameEvent) error
	GetByGameID(gameID int) ([]models.GameEvent, error)
	Delete(id int) error
}

type ScoreRepository interface {
	Create(score *models.Score) error
	GetByGameID(gameID int) ([]models.Score, error)
	GetByPlayerID(playerID int) ([]models.Score, error)
	Delete(id int) error
}

type StageRepository interface {
	Create(stage *models.Stage) error
	GetByID(id int) (*models.Stage, error)
	Update(stage *models.Stage) error
	Delete(id int) error
	GetByTournamentID(tournamentID int) ([]models.Stage, error)
}

type TournamentRepository interface {
	Create(tournament *models.Tournament) error
	GetByID(id int) (*models.Tournament, error)
	Update(tournament *models.Tournament) error
	Delete(id int) error
	List(offset, limit int) ([]models.Tournament, error)
}

type TournamentAdminRepository interface {
	Create(admin *models.TournamentAdmin) error
	GetByTournamentID(tournamentID int) ([]models.TournamentAdmin, error)
	DeleteByUserAndTournament(userID, tournamentID int) error
}

type TournamentInvitationRepository interface {
	Create(invitation *models.TournamentInvitation) error
	GetByTournamentID(tournamentID int) ([]models.TournamentInvitation, error)
	GetByClubID(clubID int) ([]models.TournamentInvitation, error)
	Delete(id int) error
}

type TournamentTeamRepository interface {
	Create(tournamentTeam *models.TournamentTeam) error
	GetByTournamentID(tournamentID int) ([]models.TournamentTeam, error)
	GetByTeamID(teamID int) ([]models.TournamentTeam, error)
	Delete(id int) error
}
