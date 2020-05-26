create view music_video_view as

  with projects_users as (
    select projects.id, projects.owner, projects.name as project_name, users.profile_photo_url, users.name as owner_name from projects
    inner join users on projects.owner = users.firebase_user_id
  )

select mv.*, p.project_name, p.owner, s.title, p.profile_photo_url as owner_photo, p.owner_name from music_videos as mv
inner join projects_users as p on mv.project = p.id
inner join songs as s on mv.song_id = s.id;
