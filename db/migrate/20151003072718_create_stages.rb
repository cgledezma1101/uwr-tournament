class CreateStages < ActiveRecord::Migration
	def change
		create_table :stages do |t|
			t.integer :tournament_id
			t.string :name

			t.timestamps
		end
	end
end
