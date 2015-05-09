class CreateClubAdmins < ActiveRecord::Migration
  def change
    create_join_table :users, :clubs, table_name: :club_admins do |t|
      t.index :user_id
      t.index :club_id
    end
  end
end
