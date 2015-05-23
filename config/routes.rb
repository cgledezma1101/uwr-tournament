UwrTournament::Application.routes.draw do
  root 'games#new'

  devise_for :users

  get '/statistics', to: 'public#statistics'

  resources :club_admins, only: [:new, :create, :destroy] do
    get 'confirm_destroy', on: :member
  end

  resources :clubs, only: [:show, :new, :create, :edit, :update]

  resources :games, only: [:create, :new, :show, :index]

  resources :invitations, only: [:new, :create] do
    get 'accept', on: :member
    get 'decline', on: :member
  end

  resources :scores, only: [:create]

  resources :teams, only: [:new, :create, :show, :destroy] do
    get 'confirm_destroy', on: :member
    get 'players', on: :member
  end

  resources :user_clubs, only: [:destroy] do
    get 'confirm_destroy', on: :member
  end
end
