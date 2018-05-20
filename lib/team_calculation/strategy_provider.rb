class TeamCalculation::StrategyProvider

	TEAM_CALCULATION_MANUAL = 0
	TEAM_CALCULATION_LEADERBOARD = 1

	# Provides a strategy to calculate a team based on the specified configuration
	#
	# @param [Integer] team_calculation The id of the team calculation strategy to use
	# @return [TeamCalculation::CalculationStrategy] An instance that can be used to calculate the team based of the specified strategy
	def self.get_strategy(team_calculation)
		if (team_calculation == TEAM_CALCULATION_MANUAL)
			return TeamCalculation::ManualCalculationStrategy.new()
		elsif (team_calculation == TEAM_CALCULATION_LEADERBOARD)
			return TeamCalculation::LeaderboardCalculationStrategy.new()
		end

		raise "Invalid team calculation strategy provided: #{team_calculation}"
	end
end