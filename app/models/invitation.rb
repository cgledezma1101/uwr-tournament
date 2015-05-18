class Invitation < ActiveRecord::Base
  belongs_to :club
  belongs_to :user

  validates :club, presence: true
  validates :user, presence: true
end
