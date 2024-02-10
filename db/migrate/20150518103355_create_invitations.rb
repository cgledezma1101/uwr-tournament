class CreateInvitations < ActiveRecord::Migration[6.0]
	def change
		create_table :invitations do |t|
			t.integer :club_id
			t.integer :user_id
			t.boolean :is_admin

			t.timestamps
		end
	end
end
