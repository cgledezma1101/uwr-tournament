class CreateTournamentTeams < ActiveRecord::Migration[6.0]
	def change
		create_table :tournament_teams do |t|
			t.integer :tournament_id
			t.integer :team_id

			t.timestamps
		end
	end
end
