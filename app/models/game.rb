class Game < ActiveRecord::Base
  STATUS_READY = 'r'
  STATUS_STARTED = 's'
  STATUS_ENDED = 'e'

  STATUSES = [STATUS_READY, STATUS_STARTED, STATUS_ENDED]

  TIED_GAME = 't'
  WINNING_COLORS = [PlayerGame::WHITE_TEAM, PlayerGame::BLUE_TEAM, TIED_GAME]

  belongs_to :stage
  belongs_to :blue_team, class_name: 'Team'
  belongs_to :white_team, class_name: 'Team'

  has_many :game_events

  has_many :scores, dependent: :destroy

  has_many :player_games, dependent: :destroy
  has_many :players, through: :player_games

  validates :blue_team, presence: true, unless: :has_ended?
  validates :white_team, presence: true, unless: :has_ended?
  validates :status, presence: true

  validate :different_teams
  validate :player_count
  validate :validate_status
  validate :correct_winning_color, if: :has_ended?

  # Determines the amount of goals that have been scored by the blue team
  #
  # @return [Integer] Amount of blue goals
  def blue_goals
    self.scores.joins{ player }
               .where{ player.team_id == my{self.blue_team_id} }
               .count
  end

  # Determines all of the players that compose the blue team
  #
  #  @return Array<Player> Blue team's players
  def blue_players
    Player.joins{ player_games }.where{ (player_games.game_id == my{self.id}) & (player_games.team_color == my{PlayerGame::BLUE_TEAM}) }
  end

  # Determines the amount of goals this player made on the match
  #
  # @param [Player] The player to be queried
  #
  # @return [Integer] The amount of goals this player has made
  def goals_for(player)
    self.scores.where{ player_id == my{player.id} }.count
  end

  # Determines whether the game represented has already finished
  #
  # @return [Boolean] Whether the game is over
  def has_ended?
    self.status == STATUS_ENDED
  end

  # Returns a text representing the two teams that will be playing
  #
  # @return String The name of the match
  def matchup_name
    "#{self.blue_team.name} #{I18n.t('general.versus_short')} #{self.white_team.name}"
  end

  # Determines the amount of goals that have been scored by the white team
  #
  # @return [Integer] Amount of blue goals
  def white_goals
    self.scores.joins{ player }
               .where{ player.team_id == my{self.white_team_id} }
               .count
  end

  # Determines all of the players that compose the white team
  #
  #  @return Array<Player> White team's players
  def white_players
    Player.joins{ player_games }.where{ (player_games.game_id == my{self.id}) & (player_games.team_color == my{PlayerGame::WHITE_TEAM}) }
  end

  private
    # Validates. when a game ends, if the winning color has been properly set
    def correct_winning_color
      if !WINNING_COLORS.include?(self.winning_color)
        self.errors.add(:winning_color, I18n.t('game.errors.wrong_winning_color'))
      end
    end

    # Validates that the blue and white teams are different
    def different_teams
      unless self.blue_team != self.white_team || self.blue_team == nil || self.white_team == nil
        errors.add(:blue_team, I18n.t('game.errors.same_team'))
      end
    end

    # Validates that a started or ended game has at least 6 players in each side.
    def player_count
      if (self.status == STATUS_STARTED || self.status == STATUS_ENDED)
        if self.blue_players.count < 6
          self.errors.add(:blue_players, I18n.t('game.errors.player_count', team: self.blue_team.name))
        end

        if self.white_players.count < 6
          self.errors.add(:white_players, I18n.t('game.errors.player_count', team: self.white_team.name))
        end
      end
    end

    # Validates that the status is always valid
    def validate_status
      if !STATUSES.include?(self.status)
        self.errors.add(:status, I18n.t('game.errors.invalid_status'))
      end
    end
end
