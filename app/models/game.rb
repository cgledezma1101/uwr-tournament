class Game < ActiveRecord::Base
  belongs_to :blue_team, class_name: 'Team'
  belongs_to :white_team, class_name: 'Team'

  has_many :scores

  has_many :player_games, inverse_of: :game
  has_many :players, through: :player_games

  validates :blue_team, presence: true
  validates :white_team, presence: true
  validates :date, presence: true

  validate :different_teams
  validate :player_validation

  # Determines the amount of goals that have been scored by the blue team
  #
  # @return [Integer] Amount of blue goals
  def blue_goals
    self.scores.joins{ player }
               .where{ player.team_id == my{self.blue_team_id} }
               .count
  end

  # Retrieves all the players that played for the blue team on this game
  #
  # @return [Array<Player>] The blue players
  def blue_players
    self.players.where{ team_id == my{self.blue_team_id} }
  end

  # Determines the amount of goals that have been scored by the white team
  #
  # @return [Integer] Amount of blue goals
  def white_goals
    self.scores.joins{ player }
               .where{ player.team_id == my{self.white_team_id} }
               .count
  end

  # Retrieves all the players that played for the white team on this game
  #
  # @return [Array<Player>] The white players
  def white_players
    self.players.where{ team_id == my{self.white_team_id} }
  end

  private
    # Validates that the blue and white teams are different
    def different_teams
      unless self.blue_team != self.white_team
        errors.add(:blue_team, 'El equipo azul es el mismo que el blanco')
      end
    end

    # Validates that the players abide by the restrictions imposed by the game
    def player_validation
      blues = self.blue_players.count
      whites = self.white_players.count
      unless (blues >= 6) &&
             (whites >= 6) &&
             (blues + whites == self.players.count)
        errors.add(:players, 'Cada equipo debe tener por lo menos 6 ' +
                             'jugadores y todos los jugadores deben ' +
                             'pertenecer a los equipos que se enfrentan')
      end
    end
end
