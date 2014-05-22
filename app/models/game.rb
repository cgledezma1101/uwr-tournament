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

  private
    # Validates that the blue and white teams are different
    def different_teams
      unless self.blue_team != self.white_team
        errors.add(:blue_team, 'El equipo azul es el mismo que el blanco')
      end
    end

    # Validates that the players abide by the restrictions imposed by the game
    def player_validation
      in_players = self.players
      unless (in_players.select{ |p| p.team_id == self.blue_team_id }
                        .count >= 6) &&
             (in_players.select{ |p| p.team_id == self.white_team_id }
                        .count >= 6) &&
             (in_players.select{ |p| (p.team_id != self.white_team_id) &&
                                  (p.team_id != self.blue_team_id) }
                        .count == 0)
        errors.add(:players, 'Cada equipo debe tener por lo menos 6 ' +
                             'jugadores y todos los jugadores deben ' +
                             'pertenecer a los equipos que se enfrentan')
      end
    end
end
