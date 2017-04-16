class CreatePlayerGames < ActiveRecord::Migration
	def change
		create_table :player_games do |t|
			t.integer :color
			t.integer :player_id
			t.integer :game_id
		end
	end
end
