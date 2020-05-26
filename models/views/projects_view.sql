CREATE VIEW projects_view AS

  WITH user_clips AS (
    SELECT c.*, u.name, u.profile_photo_url FROM clips AS c
    INNER JOIN users AS u ON c.user_id = u.firebase_user_id
  )

  SELECT
  projects.id,
  projects.name,
  projects.status,
  projects.openshot_id,
  projects.export_id,
  projects.owner,
  json_agg(json_build_object(
    'id', user_clips.id,
    'user_id', user_clips.user_id,
    'username', user_clips.name,
    'profile_photo_url', user_clips.profile_photo_url,
    'file', user_clips.file,
    'part_id', user_clips.part_id,
    'song_id', user_clips.song_id,
    'project_id', user_clips.project_id,
    'openshot_project_id', user_clips.openshot_project_id,
    'openshot_project_url', user_clips.openshot_project_url
  )) as clips,
  json_agg(
    DISTINCT(
      SELECT x FROM (
        SELECT
        songs.id,
        songs.title,
        songs.difficulty,
        songs.parts,
        songs.length_seconds,
        songs.backing_track
      ) AS x))->0
      AS song
  FROM projects
  INNER JOIN user_clips ON user_clips.project_id = projects.id
  INNER JOIN songs ON songs.id = projects.song
  GROUP BY
  projects.id,
  projects.name,
  projects.openshot_id,
  projects.export_id,
  projects.owner,
  projects.status;
