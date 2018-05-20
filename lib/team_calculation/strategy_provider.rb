class TeamCalculation::StrategyProvider

	TEAM_CALCULATION_MANUAL = "m"
	TEAM_CALCULATION_LEADERBOARD = "l"

	# Provides a strategy to calculate a team based on the specified configuration
	#
	# @param [Integer] team_calculation The id of the team calculation strategy to use
	# @return [TeamCalculation::CalculationStrategy] An instance that can be used to calculate the team based of the specified strategy
	def self.get_strategy(team_calculation)
		actual_strategy = team_calculation.nil? ? TEAM_CALCULATION_MANUAL : team_calculation
		if (actual_strategy == TEAM_CALCULATION_MANUAL.to_s())
			return TeamCalculation::ManualCalculationStrategy.new()
		else
			strategy_params = actual_strategy.split("|")
			strategy_name = strategy_params[0]

			if (strategy_name == TEAM_CALCULATION_LEADERBOARD.to_s())
				if (strategy_params.length == 2)
					return TeamCalculation::LeaderboardCalculationStrategy.new(strategy_params[1].to_i)
				elsif (strategy_params.length == 3)
					return TeamCalculation::LeaderboardCalculationStrategy.new(strategy_params[1].to_i, strategy_params[2].to_i)
				end
			end
		end

		raise "Invalid team calculation strategy provided: #{team_calculation}"
	end
end