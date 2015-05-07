class CreateUserClubs < ActiveRecord::Migration
  def change
    create_join_table :users, :clubs, table_name: :user_clubs do |t|
      t.index :user_id
      t.index :club_id
    end
  end
end
