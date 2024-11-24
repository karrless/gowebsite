-- Удаляем существующие ограничения внешних ключей
ALTER TABLE project_tech
  DROP CONSTRAINT IF EXISTS project_tech_project_id_fkey,
  DROP CONSTRAINT IF EXISTS project_tech_tech_id_fkey;

-- Добавляем новые ограничения с каскадными действиями
ALTER TABLE project_tech
  ADD CONSTRAINT project_tech_project_id_fkey
    FOREIGN KEY (project_id) REFERENCES projects(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

ALTER TABLE project_tech
  ADD CONSTRAINT project_tech_tech_id_fkey
    FOREIGN KEY (tech_id) REFERENCES techs(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE;