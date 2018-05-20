class TeamCalculation::LeaderboardCalculationStrategy < TeamCalculation::CalculationStrategy
	def initialize(leaderboard_position, stage_id = nil)
		@stage_id = stage_id
		@leaderboard_position = leaderboard_position
	end

	def calculate_team(game, team_color)
		statisticable = @stage_id.nil? ? game.stage.tournament : Stage.find(@stage_id)

		return statisticable.leaderboard[@leaderboard_position][0]
	end
end