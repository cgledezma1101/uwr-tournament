class CreateUserClubs < ActiveRecord::Migration
	def change
		create_table :user_clubs do |t|
			t.integer :user_id
			t.integer :club_id

			t.timestamps
		end
	end
end
