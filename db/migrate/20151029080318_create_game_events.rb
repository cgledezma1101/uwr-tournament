class CreateGameEvents < ActiveRecord::Migration[6.0]
	def change
		create_table :game_events do |t|
			t.string :text
			t.integer :game_id

			t.timestamps
		end
	end
end
