class Game < ActiveRecord::Base
  has_one :blue_team, class_name: 'Team'
  has_one :white_team, class_name: 'Team'

  has_many :scores

  has_many :player_games
  has_many :players, through: :player_games

  validates :blue_team, presence: true
  validates :white_team, presence: true
  validates :date, presence: true

  validate :different_teams
  validate :player_validation

  private
    # Validates that the blue and white teams are different
    def different_teams
      self.blue_team != self.white_team
    end

    # Validates that the players abide by the restrictions imposed by the game
    def player_validation
      player_amount = self.player_games
                          .where{ color == PlayerGame::BLUE}.count >= 6 &&
                      self.player_games
                          .where{ color == PlayerGame::WHITE }.count >= 6
      return false unless player_amount

      blue_team = self.player_games
                      .joins{ player }
                      .where{ (player_games.color == PlayerGame::BLUE) &
                              (player_games.player.team_id !=
                                my{self.blue_team_id}) }
      return false unless blue_team.empty?

      white_team = player_games.joins{ player }
                               .where{ (player_games.color ==
                                         PlayerGame::WHITE) &
                                       (player_games.player.team_id !=
                                         my{self.white_team_id}) }
      return false unless white_team.empty?

      true
    end
end
