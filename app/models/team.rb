class Team < ActiveRecord::Base
  has_many :players

  has_many :blue_games, class_name: 'Game', foreign_key: :blue_team_id, dependent: :destroy
  has_many :white_games, class_name: 'Game', foreign_key: :white_team_id, dependent: :destroy

  belongs_to :club

  validates :name, presence: true, uniqueness: { scope: :club }, allow_blank: false
  validates :club, presence: true

  validate :player_count

  before_destroy :dependents_set_nil

  # Returns all the games that this team has played
  #
  # @return [Array<Game>] The games played by this team
  def games
    Game.where{ (blue_team_id == my{self.id}) | (white_team_id == my{self.id}) }
  end

  # Determines how many goals the team has made
  #
  # @return [Integer] The amount of goals the team has made
  def goals
    Score.joins{ game }
         .joins{ player }
         .where{ (player.team_id == my{self.id}) &
                 ((game.blue_team_id == my{self.id}) |
                  (game.white_team_id == my{self.id})) }
         .count
  end

  # Determines how many goals have been made to this team
  #
  # @return [Integer] The amount of goals that have been made to the team
  def goals_against
    Score.joins{ game }
         .joins{ player }
         .where{ (player.team_id != my{self.id}) &
                 ((game.blue_team_id == my{self.id}) |
                  (game.white_team_id == my{self.id})) }
         .count
  end

  # Determines how many points this team has in the tournament
  #
  # @return [Integer] The amount of points the team has
  def points
    points_won = self.won_games.size * 3
    points_tied = self.tied_games.size * 1
    points_won + points_tied
  end

  # Returns the games the team has tied
  #
  # @return [Array<Game>] The games tied
  def tied_games
    Game.includes(:scores)
        .where{ (blue_team_id == my{self.id}) | (white_team_id == my{self.id}) }
        .select do |game|
          (game.blue_team_id == self.id || game.white_team_id == self.id) &&
          (game.blue_goals == game.white_goals)
        end
  end

  # Returns the games the team has won
  #
  # @return [Array<Game>] The games won
  def won_games
    Game.includes(:scores)
        .where{ (blue_team_id == my{self.id}) | (white_team_id == my{self.id}) }
        .select do |game|
          (game.blue_team_id == self.id &&
           game.blue_goals > game.white_goals) ||
          (game.white_team_id == self.id &&
           game.white_goals > game.blue_goals)
        end
  end

  # Retrieves the User objects corresponding to the players of this team
  #
  # @return [Array<User>] User information of the team's players
  def active_users
    User.joins{ players }.where{ (players.team_id == my{self.id}) & (players.is_active == true) }.uniq.to_a
  end

  # Retrieves the user information of all the members of the club this team belongs to
  #
  # @return [Array<User>] User information of the members of this team's club
  def club_members
    self.club.members
  end

  # Retrieves the players that are currently participating in the team
  #
  # @return [Array<Player>] Currently active players
  def active_players
    self.players.where{ is_active == true }.order(:number)
  end

  private

  # Sets all of the player's associations to reference a 'nil' team, so they won't reference a dead record
  def dependents_set_nil
    self.players.update_all(team_id: nil)
    self.blue_games.update_all(blue_team_id: nil)
    self.white_games.update_all(white_team_id: nil)
  end

  # Validates that the team has less than 15 players
  def player_count
    if self.players.count > 15
      self.errors.add(:players, I18n.t('team.errors.too_many_players'))
    end
  end
end
