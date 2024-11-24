-- Удаляем каскадные ограничения
ALTER TABLE project_tech
  DROP CONSTRAINT IF EXISTS project_tech_project_id_fkey,
  DROP CONSTRAINT IF EXISTS project_tech_tech_id_fkey;

-- Восстанавливаем старые ограничения без каскада (если они были, например, только на `ON DELETE RESTRICT`)
ALTER TABLE project_tech
  ADD CONSTRAINT project_tech_project_id_fkey
    FOREIGN KEY (project_id) REFERENCES projects(id);

ALTER TABLE project_tech
  ADD CONSTRAINT project_tech_tech_id_fkey
    FOREIGN KEY (tech_id) REFERENCES techs(id);