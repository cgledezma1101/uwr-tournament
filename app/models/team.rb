class Team < ActiveRecord::Base
  has_many :players

  has_many :white_games, class_name: 'Game', inverse_of: :white_team
  has_many :blue_games, class_name: 'Game', inverse_of: :blue_team

  validates :name, presence: true
end
