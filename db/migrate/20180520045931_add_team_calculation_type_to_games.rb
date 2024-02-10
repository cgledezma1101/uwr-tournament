class AddTeamCalculationTypeToGames < ActiveRecord::Migration[6.0]
  def change
    remove_column :games, :game_type
    add_column :games, :blue_team_calculation, :integer
    add_column :games, :white_team_calculation, :integer
  end
end
