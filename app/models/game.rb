class Game < ActiveRecord::Base
  belongs_to :blue_team, class_name: 'Team'
  belongs_to :white_team, class_name: 'Team'

  has_many :scores, dependent: :destroy

  has_many :blue_player_games, class_name: 'PlayerGame', inverse_of: :game, dependent: :destroy
  has_many :blue_players, through: :blue_player_games, source: :player

  has_many :white_player_games, class_name: 'PlayerGame', inverse_of: :game, dependent: :destroy
  has_many :white_players, through: :white_player_games, source: :player

  validates :blue_team, presence: true unless self.has_ended?
  validates :white_team, presence: true unless self.has_ended?
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

  # Determines the amount of goals this player made on the match
  #
  # @param [Player] The player to be queried
  #
  # @return [Integer] The amount of goals this player has made
  def goals_for(player)
    self.scores.where{ player_id == my{player.id} }.count
  end

  # Determines the amount of goals that have been scored by the white team
  #
  # @return [Integer] Amount of blue goals
  def white_goals
    self.scores.joins{ player }
               .where{ player.team_id == my{self.white_team_id} }
               .count
  end

  private
    # Validates that the blue and white teams are different
    def different_teams
      unless self.blue_team != self.white_team
        errors.add(:blue_team, I18n.t('game.errors.same_team'))
      end
    end

    # Validates that the players abide by the restrictions imposed by the game
    def player_validation
      blues = self.blue_players.count
      whites = self.white_players.count

      validate_players_in_team(self.blue_players, self.blue_team)
      validate_players_in_team(self.white_players, self.white_team)

      unless (blues >= 6) &&
             (whites >= 6) &&
             (blues + whites == self.players.size) &&
        errors.add(:players, I18n.t('game.errors.players'))
      end
    end

    # Validates whether the players in the given list belong to the specified team. Adds an error for every instance that fails
    def validate_players_in_team(players, team)
      for player in players
        if team.players.find_by(id: blue.id).nil?
          errors.add(:players, I18n.t('game.errors.player_not_in_team', player_name: blue.name, team_name: team.name))
        end
      end
    end

    # Determines whether the game represented has already finished
    #
    # @return [Boolean] Whether the game is over
    def has_ended?
      self.winning_color != nil
    end
end
