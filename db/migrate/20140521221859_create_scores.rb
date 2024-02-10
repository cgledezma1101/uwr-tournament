class CreateScores < ActiveRecord::Migration[6.0]
	def change
		create_table :scores do |t|
			t.integer :player_id
			t.integer :game_id

			t.timestamps
		end
	end
end
