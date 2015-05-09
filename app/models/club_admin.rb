class ClubAdmin < ActiveRecord::Base
  belongs_to :user
  belongs_to :club

  validate :user, presence: true
  validate :club, presence: true
end
