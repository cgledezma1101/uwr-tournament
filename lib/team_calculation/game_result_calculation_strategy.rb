class TeamCalculation::GameResultCalculationStrategy < TeamCalculation::CalculationStrategy
	WINNER = 'w'
	LOSER = 'l'

	def initialize(game_id, outcome)
		@game_id = game_id
		@outcome = outcome
	end

	def calculate_team(game, team_color)
		required_game = Game.find(@game_id)

		return retrieve_team(game, @outcome == WINNER ? PlayerGame::BLUE_TEAM : PlayerGame::WHITE_TEAM) if required_game.winning_color.nil?

		winning_color = @game.winning_color
		required_color = @outcome == WINNER ?
			winning_color :
			winning_color == PlayerGame::BLUE_TEAM ? PlayerGame::WHITE_TEAM : PlayerGame::BLUE_TEAM

		return retrieve_team(game, required_color)
	end

	private

	def retrieve_team(game, team_color)
		if team_color == PlayerGame::BLUE_TEAM
			return game.blue_team
		else
			return game.white_team
		end
	end
end