class Game < ActiveRecord::Base
  has_one :blue_team, class_name: 'Team'
  has_one :white_team, class_name: 'Team'

  has_many :blue_players, class_name: 'Player'
  has_many :white_players, class_name: 'Player'

  validates :blue_team, presence: true
  validates :white_team, presence: true
  validates :date, presence: true

  validate :player_validation

  private
    # Validates that the players abide by the restrictions imposed by the game
    def player_validation
      player_amount = self.blue_players.count >= 6 &&
                      self.white_players.count >= 6
      return false unless player_amount

      blue_team = blue_players.where{ team_id != my{self.blue_team_id} }
      return false unless blue_team.empty?

      white_team = white_players.where{ team_id != my{self.white_team_id} }
      return false unless white_team.empty?

      true
    end
end
