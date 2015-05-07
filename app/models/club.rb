class Club < ActiveRecord::Base
  has_many :user_clubs, dependent: :destroy
  has_many :users, through: :user_clubs

  has_many :teams, dependent: :destroy
end
