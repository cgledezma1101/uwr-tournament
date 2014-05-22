class Team < ActiveRecord::Base
  has_many :players

  has_many :blue_games, class_name: 'Game', foreign_key: :blue_team_id
  has_many :white_games, class_name: 'Game', foreign_key: :white_team_id

  validates :name, presence: true
end
