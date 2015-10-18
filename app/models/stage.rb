class Stage < ActiveRecord::Base
  belongs_to :tournament
  has_many :games

  validates :tournament, presence: true
  validates :name, presence: true, uniqueness: { scope: :tournament }
  validate :name_not_blank

  # Determines whether any of the games contained in this stage has ended
  #
  # @return [Boolean] Whether there's a finished game in this stage
  def has_ended_games?
    self.games.to_a.count{ |game| game.has_ended? } != 0
  end

  #########################################
  ####### STATISTICABLE MEMBERS ##########
  #########################################
  def goals_average(team)
    return self.goals_scored(team) - self.goals_received(team)
  end

  def goals_received(team)
    return Score.joins{ player }
               .joins{ game }
               .where{
                  (game.stage_id == my{self.id}) &
                  (game.status == my{Game::STATUS_ENDED}) &
                  ((game.blue_team_id == my{team.id}) | (game.white_team_id == my{team.id})) &
                  (player.team_id != my{team.id}) }
                .count
  end

  def goals_scored(team)
    return Score.joins{ player }
                .joins{ game }
                .where{
                  (game.stage_id == my{self.id}) &
                  (game.status == my{Game::STATUS_ENDED}) &
                  ((game.blue_team_id == my{team.id}) | (game.white_team_id == my{team.id})) &
                  (player.team_id == my{team.id}) }
                .count
  end

  def least_defeated
    teams = (Team.joins{ blue_games }.where{ (blue_games.stage_id == my{self.id}) & (game.status == my{Game::STATUS_ENDED}) }.to_a +
             Team.joins{ white_games }.where{ (white_games.stage_id == my{self.id}) & (game.status == my{Game::STATUS_ENDED}) }.to_a).uniq
    return teams.sort{ |team0, team1| self.goals_received(team0) <=> self.goals_received(team1) }
  end

  def lost_games(team)
    return games.where{ ((white_team_id == my{team.id}) & (winning_color == my{PlayerGame::BLUE_TEAM})) |
                        ((blue_team_id == my{team.id}) & (winning_color == my{PlayerGame::WHITE_TEAM})) }
               .count
  end

  def participating_teams
    teams = (Team.joins{ blue_games }.where{ blue_games.stage_id == my{self.id} }.to_a +
             Team.joins{ white_games }.where{ white_games.stage_id == my{self.id} }.to_a).uniq

    teams_points = teams.collect{ |team| [team, self.points_for(team)] }.to_h

    return teams.sort do |team0, team1|
      compare = 0
      if teams_points[team0] < teams_points[team1]
        compare = 1
      elsif teams_points[team0] > teams_points[team1]
        compare = -1
      else
        average_0 = self.goals_average(team0)
        average_1 = self.goals_average(team1)

        if average_0 < average_1
          compare = 1
        elsif average_0 > average_1
          compare = -11
        end
      end

      compare
    end
  end

  def points_for(team)
    return 3 * self.won_games(team) + self.tied_games(team)
  end

  def tied_games(team)
    return games.where{ ((white_team_id == my{team.id}) | (blue_team_id == my{team.id})) & (winning_color == my{Game::TIED_GAME}) }
                .count
  end

  def top_scorers
    players = Player.joins{ games }.where{ (games.stage_id == my{self.id}) & (games.status == my{Game::STATUS_ENDED}) }.uniq

    players_points = (players.map do |player|
      scores = self.games.to_a.inject(0){ |sum, game| sum + game.goals_for(player) }
      [player, scores]
    end).to_h

    players.sort{ |player0, player1| players_points[player0] < players_points[player1] ? 1 : -1 }.map{ |player| [player, players_points[player]] }
  end

  def won_games(team)
    return games.where{ ((white_team_id == my{team.id}) & (winning_color == my{PlayerGame::WHITE_TEAM})) |
                        ((blue_team_id == my{team.id}) & (winning_color == my{PlayerGame::BLUE_TEAM})) }
               .count
  end

  #########################################
  #### END OF STATISTICABLE MEMBERS #######
  #########################################

  private

  # Validates that the name that has been assigned to this stage is not null or empty
  def name_not_blank
    if self.name.blank?
      self.errors.add(:name, I18n.t('stage.errors.blank_name'))
    end
  end
end
