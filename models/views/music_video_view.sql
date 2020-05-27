create view music_video_view as

  WITH user_clips AS (
    SELECT c.*, u.name, u.profile_photo_url, sp.type  FROM clips AS c
    INNER JOIN users AS u ON c.user_id = u.firebase_user_id
    inner join song_parts as sp on c.part_id = sp.id

  ),

  projects_users as (
    select projects.id, projects.owner, projects.name as project_name, users.profile_photo_url, users.name as owner_name from projects
    inner join users on projects.owner = users.firebase_user_id
  )

select mv.id, mv.url, mv.created, mv.song_id, mv.status, mv.project, p.project_name, p.owner, s.title, p.profile_photo_url as owner_photo, p.owner_name, s.album_art,
json_agg(json_build_object(
  'username', user_clips.name,
  'profile_photo_url', user_clips.profile_photo_url,
  'part_id', user_clips.part_id,
  'part_type', user_clips.type
)) as clips
from music_videos as mv
inner join projects_users as p on mv.project = p.id
inner join songs as s on mv.song_id = s.id
INNER JOIN user_clips ON user_clips.project_id = mv.project
group by mv.id, mv.url, mv.created, mv.song_id, mv.status, mv.project,
p.project_name, p.owner, s.title, p.profile_photo_url, p.owner_name, s.album_art;
