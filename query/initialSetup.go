package query

const CREATE_TABLE_MODALITIES = `

	CREATE TABLE IF NOT EXISTS public.modalities (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL
	);
`
const CREATE_TABLE_REGISTRATION = `CREATE TABLE IF NOT EXISTS public.registration (
	id SERIAL PRIMARY KEY,
	modalities TEXT [],
	day_to_registry DATE,
	status BOOLEAN
);`

const CREATE_TABLE_PLAN = `CREATE TABLE IF NOT EXISTS public.plan (
	id SERIAL PRIMARY KEY,
	price DECIMAL NOT NULL,
	workouts INTEGER NOT NULL
);`

const CREATE_TABLE_STUDY = `CREATE TABLE IF NOT EXISTS public.study (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	rg TEXT NOT NULL,
	registry_id INTEGER NOT NULL,
	plan_id INTEGER NOT NULL,
	CONSTRAINT fk_registry FOREIGN KEY (registry_id) REFERENCES registration (id),
	CONSTRAINT fk_plan FOREIGN KEY (plan_id) REFERENCES plan (id)
);`

const CREATE_TABLE_PAYMENT = `CREATE TABLE IF NOT EXISTS public.payment (
	id SERIAL PRIMARY KEY,
	due_date DATE NOT NULL,
	ref_date DATE NOT NULL,
	payment_status BOOLEAN,
	study_id INTEGER NOT NULL,
	FOREIGN KEY (study_id) REFERENCES study (id)
)`
