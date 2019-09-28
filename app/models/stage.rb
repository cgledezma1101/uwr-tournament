class Stage < ApplicationRecord
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
        return Score
            .joins(:player, :game)
            .where(games: { stage_id: self.id, status: Game::STATUS_ENDED })
            .where.not(players: { team_id: team.id })
            .where(games: { blue_team_id: team.id }).or(
                Score.joins(:player, :game).where(games: { white_team_id: team.id })
            )
            .count
    end

    def goals_scored(team)
        return Score
            .joins(:player, :game)
            .where(
                games: { stage_id: self.id, status: Game::STATUS_ENDED },
                players: { team_id: team.id }
            )
            .where(games: { blue_team_id: team.id }).or(
                Score.joins(:player, :game).where(games: { white_team_id: team.id })
            )
            .count
    end

    def least_defeated
        blue_teams = Team
            .joins(:blue_games)
            .where(games: { stage_id: self.id, status: Game::STATUS_ENDED })
            .to_a

        white_teams = Team
            .joins(:white_games)
            .where(games: { stage_id: self.id, status: Game::STATUS_ENDED })
            .to_a

        teams = (blue_teams + white_teams).uniq

        goals_against = teams.map{ |team| [team, self.goals_received(team)] }
        return goals_against.sort{ |against0, against1| against0[1] <=> against1[1] }
    end

    def lost_games(team)
        return games
            .where(white_team_id: team.id, winning_color: PlayerGame::BLUE_TEAM).or(
                games.where(blue_team_id: team.id, winning_color: PlayerGame::WHITE_TEAM)
            )
            .count
    end

    def leaderboard
        blue_teams = Team
            .joins(:blue_games)
            .where(games: { stage_id: self.id })
            .to_a

        white_teams = Team
            .joins(:white_games)
            .where(games: { stage_id: self.id })
            .to_a
        teams = (blue_teams + white_teams).uniq

        teams.collect{ |team| [team, self.points_for(team)] }
            .sort do |team0, team1|
                compare = 0
                if team0[1] < team1[1]
                    compare = 1
                elsif team0[1] > team1[1]
                    compare = -1
                else
                    average_0 = self.goals_average(team0[0])
                    average_1 = self.goals_average(team1[0])

                    if average_0 < average_1
                        compare = 1
                    elsif average_0 > average_1
                        compare = -1
                    end
                end

                compare
            end
    end

    def points_for(team)
        return 3 * self.won_games(team) + self.tied_games(team)
    end

    def tied_games(team)
        return games
            .where(winning_color: Game::TIED_GAME)
            .where(white_team_id: team.id).or(games.where(blue_team_id: team.id))
            .count
    end

    def top_scorers
        players = Player.joins(:games).where(games: { stage_id: self.id, status: Game::STATUS_ENDED }).uniq

        players_points = players.map do |player|
            scores = Score.joins(:game).where(games: { stage_id: self.id }, player_id: player.id).count
            [player, scores]
        end

        players_points.sort{ |points0, points1| points0[1] < points1[1] ? 1 : -1 }
    end

    def won_games(team)
        return games
            .where(white_team_id: team.id, winning_color: PlayerGame::WHITE_TEAM).or(
                games.where(blue_team_id: team.id, winning_color: PlayerGame::BLUE_TEAM)
            )
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
