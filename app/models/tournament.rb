class Tournament < ActiveRecord::Base
  has_many :tournament_admins, dependent: :destroy
  has_many :admins, through: :tournament_admins, source: :user

  has_many :tournament_invitations, dependent: :destroy
  has_many :clubs, through: :tournament_invitations

  has_many :tournament_teams, dependent: :destroy
  has_many :teams, through: :tournament_teams

  has_many :stages, dependent: :destroy

  validates :name, presence: true
  validates :start_date, presence: true
  validates :end_date, presence: true
  validate :date_range

  #########################################
  ####### STATISTICABLE MEMBERS ##########
  #########################################

  def goals_average(team)
    return self.goals_scored(team) - self.goals_received(team)
  end

  def goals_received(team)
    self.stages.inject(0){ |sum, stage| sum + stage.goals_received(team) }
  end

  def goals_scored(team)
    self.stages.inject(0){ |sum, stage| sum + stage.goals_scored(team) }
  end

  def least_defeated
    teams_defeats = self.teams.to_a.map{ |team| [team, self.goals_received(team)] }
    teams_defeats.sort{ |defeat0, defeat1| defeat0[1] <=> defeat1[1] }
  end

  def lost_games(team)
    self.stages.inject(0){ |sum, stage| sum + stage.lost_games(team) }
  end

  def leaderboard
    teams.to_a
         .map{ |team| [team, self.points_for(team)] }
         .sort do |team0, team1|
           compare = 0
           if team0[1] < team1[1]
             compare = 1
           elsif team0[1] > team1[1]
             compare = -1
           else
             average0 = self.goals_average(team0[0])
             average1 = self.goals_average(team1[0])

             if average0 < average1
               compare = 1
             elsif average0 > average1
               compare = -1
             end
           end

           compare
         end
  end

  def points_for(team)
    3 * self.won_games(team) + self.tied_games(team)
  end

  def tied_games(team)
    self.stages.inject(0){ |sum, stage| sum + stage.tied_games(team) }
  end

  def top_scorers
    players = Player.joins{ games.stage }.where{ (games.status == my{Game::STATUS_ENDED}) & (games.stage.tournament_id == my{self.id}) }

    player_points = players.map do |player|
      scores = Score.joins{ game.stage }.where{ (player_id == my{player.id}) & (game.stage.tournament_id == my{self.id}) }.count
      [player, scores]
    end
    player_points.sort{ |points0, points1| points0[1] < points1[1] ? 1 : -1 }
  end

  def won_games(team)
    self.stages.inject(0){ |sum, stage| sum + stage.won_games(team) }
  end

  #########################################
  #### END OF STATISTICABLE MEMBERS #######
  #########################################

  private

  # Validates that the start date of the tournament is before the end date
  def date_range
    self.errors.add(:start_date, I18n.t('tournament.date_mismatch')) if start_date > end_date
  end
end
