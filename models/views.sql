
CREATE VIEW projects_view AS

  SELECT
  projects.id,
  projects.name,
  projects.status,
  json_agg(json_build_object(
    'firebase_user_id', users.firebase_user_id,
    'id', users.id,
    'email', users.email,
    'name', users.name,
    'profile_photo_url', users.profile_photo_url
  )) as users,
  json_agg(json_build_object(
    'id', clips.id,
    'user_id', clips.user_id,
    'file', clips.file,
    'part_id', clips.part_id,
    'song_id', clips.song_id,
    'project_id', clips.project_id,
    'openshot_project_id', clips.openshot_project_id,
    'openshot_project_url', clips.openshot_project_url
  )) as clips,
  json_agg(
    DISTINCT(
      SELECT x FROM (
        SELECT
        songs.id,
        songs.title,
        songs.difficulty,
        songs.parts
      ) AS x))->0
      AS song
  FROM projects
  INNER JOIN clips ON clips.project_id = projects.id
  INNER JOIN users ON firebase_user_id = clips.user_id
  INNER JOIN songs ON songs.id = projects.song
  GROUP BY
  projects.id,
  projects.name,
  projects.status;


  WITH user_projects AS (
    SELECT p.* FROM clips AS c
    INNER JOIN projects AS p ON c.project_id = p.id
    WHERE c.user_id = '3QQFkG3sqlMjJKu5FDC1eFrFdFq1'
  )
