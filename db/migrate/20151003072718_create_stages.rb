class CreateStages < ActiveRecord::Migration[6.0]
	def change
		create_table :stages do |t|
			t.integer :tournament_id
			t.string :name

			t.timestamps
		end
	end
end
