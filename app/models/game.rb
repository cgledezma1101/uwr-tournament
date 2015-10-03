class Game < ActiveRecord::Base
  STATUS_READY = 'r'
  STATUS_STARTED = 's'
  STATUS_ENDED = 'e'

  STATUSES = [STATUS_READY, STATUS_STARTED, STATUS_ENDED]

  belongs_to :stage
  belongs_to :blue_team, class_name: 'Team'
  belongs_to :white_team, class_name: 'Team'

  has_many :scores, dependent: :destroy

  has_many :blue_player_games, class_name: 'PlayerGame', inverse_of: :game, dependent: :destroy
  has_many :blue_players, through: :blue_player_games, source: :player

  has_many :white_player_games, class_name: 'PlayerGame', inverse_of: :game, dependent: :destroy
  has_many :white_players, through: :white_player_games, source: :player

  validates :blue_team, presence: true, unless: :has_ended?
  validates :white_team, presence: true, unless: :has_ended?
  validates :status, presence: true

  validate :different_teams
  validate :validate_status

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
      unless self.blue_team != self.white_team || self.blue_team == nil || self.white_team == nil
        errors.add(:blue_team, I18n.t('game.errors.same_team'))
      end
    end

    # Determines whether the game represented has already finished
    #
    # @return [Boolean] Whether the game is over
    def has_ended?
      self.status == STATUS_ENDED
    end

    # Validates that the status is always valid
    def validate_status
      if !STATUSES.include?(self.status)
        self.errors.add(:status, I18n.t('game.errors.invalid_status'))
      end
    end
end
