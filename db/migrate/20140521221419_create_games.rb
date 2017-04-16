class CreateGames < ActiveRecord::Migration
	def change
		create_table :games do |t|
			t.date :date
			t.integer :blue_team_id
			t.integer :white_team_id

			t.timestamps
		end
	end
end
