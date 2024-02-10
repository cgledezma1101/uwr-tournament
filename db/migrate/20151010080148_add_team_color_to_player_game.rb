class AddTeamColorToPlayerGame < ActiveRecord::Migration[6.0]
	def change
		add_column :player_games, :team_color, :string
	end
end
