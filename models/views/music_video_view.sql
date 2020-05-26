create view music_video_view as
select mv.*, p.name, p.owner, s.title from music_videos as mv
inner join projects as p on mv.project = p.id
inner join songs as s on mv.song_id = s.id;
