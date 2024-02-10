class CreateClubJoinRequests < ActiveRecord::Migration[6.0]
	def change
		create_table :club_join_requests do |t|
			t.integer :user_id
			t.integer :club_id

			t.timestamps
		end
	end
end
