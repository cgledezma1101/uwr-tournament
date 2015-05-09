class CreateClubAdmins < ActiveRecord::Migration
  def change
    create_table :club_admins do |t|
      t.integer :user_id
      t.integer :club_id

      t.timestamps
    end
  end
end
