package ogame

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInt(t *testing.T) {
	assert.Equal(t, int64(1234567890), ParseInt("1.234.567.890"))
}

func TestToInt(t *testing.T) {
	assert.Equal(t, 1234567890, toInt([]byte("1234567890")))
}

func TestMinInt(t *testing.T) {
	assert.Equal(t, int64(2), MinInt(5, 2, 3))
}

func TestParseCoord(t *testing.T) {
	coord, _ := ParseCoord("[P:1:2:3]")
	assert.Equal(t, Coordinate{1, 2, 3, PlanetType}, coord)
	coord, _ = ParseCoord("[M:1:2:3]")
	assert.Equal(t, Coordinate{1, 2, 3, MoonType}, coord)
	coord, _ = ParseCoord("M:1:2:3")
	assert.Equal(t, Coordinate{1, 2, 3, MoonType}, coord)
	coord, _ = ParseCoord("1:2:3")
	assert.Equal(t, Coordinate{1, 2, 3, PlanetType}, coord)
	coord, _ = ParseCoord("1:2:3")
	assert.Equal(t, Coordinate{1, 2, 3, PlanetType}, coord)
	coord, _ = ParseCoord("D:1:2:3")
	assert.Equal(t, Coordinate{1, 2, 3, DebrisType}, coord)
	coord, _ = ParseCoord("[D:1:2:3]")
	assert.Equal(t, Coordinate{1, 2, 3, DebrisType}, coord)
	_, err := ParseCoord("[A:1:2:3]")
	assert.NotNil(t, err)
	_, err = ParseCoord("aP:1:2:3")
	assert.NotNil(t, err)
	_, err = ParseCoord("P:1234:2:3")
	assert.NotNil(t, err)
	_, err = ParseCoord("P:1:2345:3")
	assert.NotNil(t, err)
	_, err = ParseCoord("P:1:2:3456")
	assert.NotNil(t, err)
}

func TestName2id(t *testing.T) {
	assert.Equal(t, ID(0), ShipName2ID("Not valid"))
	assert.Equal(t, LightFighterID, ShipName2ID("Light Fighter"))
	assert.Equal(t, LightFighterID, ShipName2ID("Chasseur léger"))
	assert.Equal(t, LightFighterID, ShipName2ID("Leichter Jäger"))
	assert.Equal(t, LightFighterID, ShipName2ID("Caça Ligeiro"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Caça Pesado"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Großer Transporter"))
	assert.Equal(t, DestroyerID, ShipName2ID("Zerstörer"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Nave pequeña de carga"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Satélite solar"))
	assert.Equal(t, ID(0), ShipName2ID("人中位"))

	// fi
	assert.Equal(t, LightFighterID, ShipName2ID("Kevyt Hävittäjä"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Raskas Hävittäjä"))
	assert.Equal(t, CruiserID, ShipName2ID("Risteilijä"))
	assert.Equal(t, BattleshipID, ShipName2ID("Taistelualus"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Taisteluristeilijä"))
	assert.Equal(t, BomberID, ShipName2ID("Pommittaja"))
	assert.Equal(t, DestroyerID, ShipName2ID("Tuhoaja"))
	assert.Equal(t, DeathstarID, ShipName2ID("Kuolemantähti"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Pieni rahtialus"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Suuri rahtialus"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Siirtokunta-alus"))
	assert.Equal(t, RecyclerID, ShipName2ID("Kierrättäjä"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Vakoiluluotain"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Aurinkosatelliitti"))
	assert.Equal(t, CrawlerID, ShipName2ID("Crawler"))
	assert.Equal(t, ReaperID, ShipName2ID("Reaper"))
	assert.Equal(t, PathfinderID, ShipName2ID("Pathfinder"))

	// hu
	assert.Equal(t, LightFighterID, ShipName2ID("Könnyű Harcos"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Nehéz Harcos"))
	assert.Equal(t, CruiserID, ShipName2ID("Cirkáló"))
	assert.Equal(t, BattleshipID, ShipName2ID("Csatahajó"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Csatacirkáló"))
	assert.Equal(t, BomberID, ShipName2ID("Bombázó"))
	assert.Equal(t, DestroyerID, ShipName2ID("Romboló"))
	assert.Equal(t, DeathstarID, ShipName2ID("Halálcsillag"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Kis szállító"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Nagy Szállító"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Kolóniahajó"))
	assert.Equal(t, RecyclerID, ShipName2ID("Szemetesek"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Kémszonda"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Napműhold"))
	assert.Equal(t, CrawlerID, ShipName2ID("Crawler"))
	assert.Equal(t, ReaperID, ShipName2ID("Kaszás"))
	assert.Equal(t, PathfinderID, ShipName2ID("Felderítő"))

	// ro
	assert.Equal(t, LightFighterID, ShipName2ID("Vanator usor"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Vanator greu"))
	assert.Equal(t, CruiserID, ShipName2ID("Crucisator"))
	assert.Equal(t, BattleshipID, ShipName2ID("Nava de razboi"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Interceptor"))
	assert.Equal(t, BomberID, ShipName2ID("Bombardier"))
	assert.Equal(t, DestroyerID, ShipName2ID("Distrugator"))
	assert.Equal(t, DeathstarID, ShipName2ID("RIP"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Transportor mic"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Transportor mare"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Nava de Colonizare"))
	assert.Equal(t, RecyclerID, ShipName2ID("Reciclator"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Proba de spionaj"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Satelit solar"))
	assert.Equal(t, CrawlerID, ShipName2ID("Crawler"))
	assert.Equal(t, ReaperID, ShipName2ID("Reaper"))
	assert.Equal(t, PathfinderID, ShipName2ID("Pathfinder"))

	// cz
	assert.Equal(t, LightFighterID, ShipName2ID("Lehký stíhač"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Těžký stíhač"))
	assert.Equal(t, CruiserID, ShipName2ID("Křižník"))
	assert.Equal(t, BattleshipID, ShipName2ID("Bitevní loď"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Bitevní křižník"))
	assert.Equal(t, BomberID, ShipName2ID("Bombardér"))
	assert.Equal(t, DestroyerID, ShipName2ID("Ničitel"))
	assert.Equal(t, DeathstarID, ShipName2ID("Hvězda smrti"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Malý transportér"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Velký transportér"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Kolonizační loď"))
	assert.Equal(t, RecyclerID, ShipName2ID("Recyklátor"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Špionážní sonda"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Solární satelit"))
	assert.Equal(t, CrawlerID, ShipName2ID("Crawler"))
	assert.Equal(t, ReaperID, ShipName2ID("Rozparovač"))
	assert.Equal(t, PathfinderID, ShipName2ID("Průzkumník"))

	// mx
	assert.Equal(t, LightFighterID, ShipName2ID("Cazador ligero"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Cazador pesado"))
	assert.Equal(t, CruiserID, ShipName2ID("Crucero"))
	assert.Equal(t, BattleshipID, ShipName2ID("Nave de batalla"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Acorazado"))
	assert.Equal(t, BomberID, ShipName2ID("Bombardero"))
	assert.Equal(t, DestroyerID, ShipName2ID("Destructor"))
	assert.Equal(t, DeathstarID, ShipName2ID("Estrella de la muerte"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Nave pequeña de carga"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Nave grande de carga"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Nave de la colonia"))
	assert.Equal(t, RecyclerID, ShipName2ID("Reciclador"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Sonda de espionaje"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Satélite solar"))
	assert.Equal(t, CrawlerID, ShipName2ID("Taladrador"))
	assert.Equal(t, ReaperID, ShipName2ID("Segador"))
	assert.Equal(t, PathfinderID, ShipName2ID("Explorador"))

	// hr
	assert.Equal(t, LightFighterID, ShipName2ID("Mali lovac"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Veliki lovac"))
	assert.Equal(t, CruiserID, ShipName2ID("Krstarica"))
	assert.Equal(t, BattleshipID, ShipName2ID("Borbeni brod"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Oklopna krstarica"))
	assert.Equal(t, BomberID, ShipName2ID("Bombarder"))
	assert.Equal(t, DestroyerID, ShipName2ID("Razarač"))
	assert.Equal(t, DeathstarID, ShipName2ID("Zvijezda smrti"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Mali transporter"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Veliki transporter"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Kolonijalni brod"))
	assert.Equal(t, RecyclerID, ShipName2ID("Recikler"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Sonde za špijunažu"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Solarni satelit"))
	assert.Equal(t, CrawlerID, ShipName2ID("Puzavac"))
	assert.Equal(t, ReaperID, ShipName2ID("Žetelac"))
	assert.Equal(t, PathfinderID, ShipName2ID("Krčilac"))

	// no
	assert.Equal(t, LightFighterID, ShipName2ID("Lett Jeger"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Tung Jeger"))
	assert.Equal(t, CruiserID, ShipName2ID("Krysser"))
	assert.Equal(t, BattleshipID, ShipName2ID("Slagskip"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Slagkrysser"))
	assert.Equal(t, BomberID, ShipName2ID("Bomber"))
	assert.Equal(t, DestroyerID, ShipName2ID("Destroyer"))
	assert.Equal(t, DeathstarID, ShipName2ID("Døds stjerne"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Lite Lasteskip"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Stort Lasteskip"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Koloni Skip"))
	assert.Equal(t, RecyclerID, ShipName2ID("Resirkulerer"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Spionasjesonde"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Solar Satelitt"))
	assert.Equal(t, CrawlerID, ShipName2ID("Crawler"))
	assert.Equal(t, ReaperID, ShipName2ID("Reaper"))
	assert.Equal(t, PathfinderID, ShipName2ID("Pathfinder"))

	// it
	assert.Equal(t, LightFighterID, ShipName2ID("Caccia Leggero"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Caccia Pesante"))
	assert.Equal(t, CruiserID, ShipName2ID("Incrociatore"))
	assert.Equal(t, BattleshipID, ShipName2ID("Nave da battaglia"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Incrociatore da Battaglia"))
	assert.Equal(t, BomberID, ShipName2ID("Bombardiere"))
	assert.Equal(t, DestroyerID, ShipName2ID("Corazzata"))
	assert.Equal(t, DeathstarID, ShipName2ID("Morte Nera"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Cargo leggero"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Cargo Pesante"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Colonizzatrice"))
	assert.Equal(t, RecyclerID, ShipName2ID("Riciclatrici"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Sonda spia"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Satellite Solare"))
	assert.Equal(t, CrawlerID, ShipName2ID("Crawler"))
	assert.Equal(t, ReaperID, ShipName2ID("Reaper"))
	assert.Equal(t, PathfinderID, ShipName2ID("Pathfinder"))

	// pl
	assert.Equal(t, LightFighterID, ShipName2ID("Lekki myśliwiec"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Ciężki myśliwiec"))
	assert.Equal(t, CruiserID, ShipName2ID("Krążownik"))
	assert.Equal(t, BattleshipID, ShipName2ID("Okręt wojenny"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Pancernik"))
	assert.Equal(t, BomberID, ShipName2ID("Bombowiec"))
	assert.Equal(t, DestroyerID, ShipName2ID("Niszczyciel"))
	assert.Equal(t, DeathstarID, ShipName2ID("Gwiazda Śmierci"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Mały transporter"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Duży transporter"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Statek kolonizacyjny"))
	assert.Equal(t, RecyclerID, ShipName2ID("Recykler"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Sonda szpiegowska"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Satelita słoneczny"))
	assert.Equal(t, CrawlerID, ShipName2ID("Pełzacz"))
	assert.Equal(t, ReaperID, ShipName2ID("Rozpruwacz"))
	assert.Equal(t, PathfinderID, ShipName2ID("Pionier"))

	// tr
	assert.Equal(t, LightFighterID, ShipName2ID("Hafif Avcı"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Ağır Avcı"))
	assert.Equal(t, CruiserID, ShipName2ID("Kruvazör"))
	assert.Equal(t, BattleshipID, ShipName2ID("Komuta Gemisi"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Firkateyn"))
	assert.Equal(t, BomberID, ShipName2ID("Bombardıman Gemisi"))
	assert.Equal(t, DestroyerID, ShipName2ID("Muhrip"))
	assert.Equal(t, DeathstarID, ShipName2ID("Ölüm Yildizi"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Küçük Nakliye Gemisi"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Büyük Nakliye Gemisi"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Koloni Gemisi"))
	assert.Equal(t, RecyclerID, ShipName2ID("Geri Dönüsümcü"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Casus Sondasi"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Solar Uydu"))
	assert.Equal(t, CrawlerID, ShipName2ID("Paletli"))
	assert.Equal(t, ReaperID, ShipName2ID("Azrail"))
	assert.Equal(t, PathfinderID, ShipName2ID("Rehber"))

	// ar
	assert.Equal(t, LightFighterID, ShipName2ID("Cazador ligero"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Cazador pesado"))
	assert.Equal(t, CruiserID, ShipName2ID("Crucero"))
	assert.Equal(t, BattleshipID, ShipName2ID("Nave de batalla"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Acorazado"))
	assert.Equal(t, BomberID, ShipName2ID("Bombardero"))
	assert.Equal(t, DestroyerID, ShipName2ID("Destructor"))
	assert.Equal(t, DeathstarID, ShipName2ID("Estrella de la muerte"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Nave pequeña de carga"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Nave grande de carga"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Nave colonizadora"))
	assert.Equal(t, RecyclerID, ShipName2ID("Reciclador"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Sonda de espionaje"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Satélite solar"))
	assert.Equal(t, CrawlerID, ShipName2ID("Taladrador"))
	assert.Equal(t, ReaperID, ShipName2ID("Segador"))
	assert.Equal(t, PathfinderID, ShipName2ID("Explorador"))

	// pt
	assert.Equal(t, LightFighterID, ShipName2ID("Caça Ligeiro"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Caça Pesado"))
	assert.Equal(t, CruiserID, ShipName2ID("Cruzador"))
	assert.Equal(t, BattleshipID, ShipName2ID("Nave de Batalha"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Interceptor"))
	assert.Equal(t, BomberID, ShipName2ID("Bombardeiro"))
	assert.Equal(t, DestroyerID, ShipName2ID("Destruidor"))
	assert.Equal(t, DeathstarID, ShipName2ID("Estrela da Morte"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Cargueiro Pequeno"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Cargueiro Grande"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Nave de Colonização"))
	assert.Equal(t, RecyclerID, ShipName2ID("Reciclador"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Sonda de Espionagem"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Satélite Solar"))
	assert.Equal(t, CrawlerID, ShipName2ID("Rastejador"))
	assert.Equal(t, ReaperID, ShipName2ID("Ceifeira"))
	assert.Equal(t, PathfinderID, ShipName2ID("Exploradora"))

	// nl
	assert.Equal(t, LightFighterID, ShipName2ID("Licht gevechtsschip"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Zwaar gevechtsschip"))
	assert.Equal(t, CruiserID, ShipName2ID("Kruiser"))
	assert.Equal(t, BattleshipID, ShipName2ID("Slagschip"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Interceptor"))
	assert.Equal(t, BomberID, ShipName2ID("Bommenwerper"))
	assert.Equal(t, DestroyerID, ShipName2ID("Vernietiger"))
	assert.Equal(t, DeathstarID, ShipName2ID("Ster des Doods"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Klein vrachtschip"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Groot vrachtschip"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Kolonisatieschip"))
	assert.Equal(t, RecyclerID, ShipName2ID("Recycler"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Spionagesonde"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Zonne-energiesatelliet"))
	assert.Equal(t, CrawlerID, ShipName2ID("Processer"))
	assert.Equal(t, ReaperID, ShipName2ID("Ruimer"))
	assert.Equal(t, PathfinderID, ShipName2ID("Navigator"))

	// dk
	assert.Equal(t, LightFighterID, ShipName2ID("Lille Jæger"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Stor Jæger"))
	assert.Equal(t, CruiserID, ShipName2ID("Krydser"))
	assert.Equal(t, BattleshipID, ShipName2ID("Slagskib"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Interceptor"))
	assert.Equal(t, BomberID, ShipName2ID("Bomber"))
	assert.Equal(t, DestroyerID, ShipName2ID("Destroyer"))
	assert.Equal(t, DeathstarID, ShipName2ID("Dødsstjerne"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Lille Transporter"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Stor Transporter"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Koloniskib"))
	assert.Equal(t, RecyclerID, ShipName2ID("Recycler"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Spionagesonde"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Solarsatellit"))
	assert.Equal(t, CrawlerID, ShipName2ID("Kravler"))
	assert.Equal(t, ReaperID, ShipName2ID("Reaper"))
	assert.Equal(t, PathfinderID, ShipName2ID("Stifinder"))

	// ru
	assert.Equal(t, LightFighterID, ShipName2ID("Лёгкий истребитель"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Тяжёлый истребитель"))
	assert.Equal(t, CruiserID, ShipName2ID("Крейсер"))
	assert.Equal(t, BattleshipID, ShipName2ID("Линкор"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Линейный крейсер"))
	assert.Equal(t, BomberID, ShipName2ID("Бомбардировщик"))
	assert.Equal(t, DestroyerID, ShipName2ID("Уничтожитель"))
	assert.Equal(t, DeathstarID, ShipName2ID("Звезда смерти"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Малый транспорт"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Большой транспорт"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Колонизатор"))
	assert.Equal(t, RecyclerID, ShipName2ID("Переработчик"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Шпионский зонд"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Солнечный спутник"))
	assert.Equal(t, CrawlerID, ShipName2ID("Гусеничник"))
	assert.Equal(t, ReaperID, ShipName2ID("Жнец"))
	assert.Equal(t, PathfinderID, ShipName2ID("Первопроходец"))

	// gr
	assert.Equal(t, LightFighterID, ShipName2ID("Ελαφρύ Μαχητικό"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Βαρύ Μαχητικό"))
	assert.Equal(t, CruiserID, ShipName2ID("Καταδιωκτικό"))
	assert.Equal(t, BattleshipID, ShipName2ID("Καταδρομικό"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Θωρηκτό Αναχαίτισης"))
	assert.Equal(t, BomberID, ShipName2ID("Βομβαρδιστικό"))
	assert.Equal(t, DestroyerID, ShipName2ID("Destroyer"))
	assert.Equal(t, DeathstarID, ShipName2ID("Deathstar"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Μικρό Μεταγωγικό"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Μεγάλο Μεταγωγικό"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Σκάφος Αποικιοποίησης"))
	assert.Equal(t, RecyclerID, ShipName2ID("Ανακυκλωτής"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Κατασκοπευτικό Στέλεχος"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Ηλιακοί Συλλέκτες"))
	assert.Equal(t, CrawlerID, ShipName2ID("Crawler"))
	assert.Equal(t, ReaperID, ShipName2ID("Reaper"))
	assert.Equal(t, PathfinderID, ShipName2ID("Pathfinder"))

	// jp
	assert.Equal(t, LightFighterID, ShipName2ID("軽戦闘機"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("重戦闘機"))
	assert.Equal(t, CruiserID, ShipName2ID("巡洋艦"))
	assert.Equal(t, BattleshipID, ShipName2ID("バトルシップ"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("大型戦艦"))
	assert.Equal(t, BomberID, ShipName2ID("爆撃機"))
	assert.Equal(t, DestroyerID, ShipName2ID("デストロイヤー"))
	assert.Equal(t, DeathstarID, ShipName2ID("デススター"))
	assert.Equal(t, SmallCargoID, ShipName2ID("小型輸送機"))
	assert.Equal(t, LargeCargoID, ShipName2ID("大型輸送機"))
	assert.Equal(t, ColonyShipID, ShipName2ID("コロニーシップ"))
	assert.Equal(t, RecyclerID, ShipName2ID("残骸回収船"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("偵察機"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("ソーラーサテライト"))
	assert.Equal(t, CrawlerID, ShipName2ID("クローラー"))
	assert.Equal(t, ReaperID, ShipName2ID("リーパー"))
	assert.Equal(t, PathfinderID, ShipName2ID("パスファインダー"))

	// sk
	assert.Equal(t, LightFighterID, ShipName2ID("Ľahký stíhač"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Ťažký stíhač"))
	assert.Equal(t, CruiserID, ShipName2ID("Krížnik"))
	assert.Equal(t, BattleshipID, ShipName2ID("Bojová loď"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Bojový krížnik"))
	assert.Equal(t, BomberID, ShipName2ID("Bombardér"))
	assert.Equal(t, DestroyerID, ShipName2ID("Devastátor"))
	assert.Equal(t, DeathstarID, ShipName2ID("Hviezda smrti"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Malý transportér"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Veľký transportér"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Kolonizačná loď"))
	assert.Equal(t, RecyclerID, ShipName2ID("Recyklátor"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Špionážna sonda"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Solárny satelit"))
	assert.Equal(t, CrawlerID, ShipName2ID("Vrták"))
	assert.Equal(t, ReaperID, ShipName2ID("Kosa"))
	assert.Equal(t, PathfinderID, ShipName2ID("Prieskumník"))

	// si
	assert.Equal(t, LightFighterID, ShipName2ID("Lahek lovec"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Težki lovec"))
	assert.Equal(t, CruiserID, ShipName2ID("Križarka"))
	assert.Equal(t, BattleshipID, ShipName2ID("Bojna ladja"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Bojna križarka"))
	assert.Equal(t, BomberID, ShipName2ID("Bombnik"))
	assert.Equal(t, DestroyerID, ShipName2ID("Uničevalec"))
	assert.Equal(t, DeathstarID, ShipName2ID("Zvezda smrti"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Majhna tovorna ladja"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Velika tovorna ladja"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Kolonizacijska ladja"))
	assert.Equal(t, RecyclerID, ShipName2ID("Recikler"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Vohunska sonda"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Sončni satelit"))
	assert.Equal(t, CrawlerID, ShipName2ID("Plazilec"))
	assert.Equal(t, ReaperID, ShipName2ID("Kombajn"))
	assert.Equal(t, PathfinderID, ShipName2ID("Iskalec sledi"))

	// fr
	assert.Equal(t, LightFighterID, ShipName2ID("Chasseur léger"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("Chasseur lourd"))
	assert.Equal(t, CruiserID, ShipName2ID("Croiseur"))
	assert.Equal(t, BattleshipID, ShipName2ID("Vaisseau de bataille"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("Traqueur"))
	assert.Equal(t, BomberID, ShipName2ID("Bombardier"))
	assert.Equal(t, DestroyerID, ShipName2ID("Destructeur"))
	assert.Equal(t, DeathstarID, ShipName2ID("Étoile de la mort"))
	assert.Equal(t, SmallCargoID, ShipName2ID("Petit transporteur"))
	assert.Equal(t, LargeCargoID, ShipName2ID("Grand transporteur"))
	assert.Equal(t, ColonyShipID, ShipName2ID("Vaisseau de colonisation"))
	assert.Equal(t, RecyclerID, ShipName2ID("Recycleur"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("Sonde d`espionnage"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("Satellite solaire"))
	assert.Equal(t, CrawlerID, ShipName2ID("Foreuse"))
	assert.Equal(t, ReaperID, ShipName2ID("Faucheur"))
	assert.Equal(t, PathfinderID, ShipName2ID("Éclaireur"))

	// tw
	assert.Equal(t, LightFighterID, ShipName2ID("輕型戰鬥機"))
	assert.Equal(t, HeavyFighterID, ShipName2ID("重型戰鬥機"))
	assert.Equal(t, CruiserID, ShipName2ID("巡洋艦"))
	assert.Equal(t, BattleshipID, ShipName2ID("戰列艦"))
	assert.Equal(t, BattlecruiserID, ShipName2ID("戰鬥巡洋艦"))
	assert.Equal(t, BomberID, ShipName2ID("導彈艦"))
	assert.Equal(t, DestroyerID, ShipName2ID("毀滅者"))
	assert.Equal(t, DeathstarID, ShipName2ID("死星"))
	assert.Equal(t, SmallCargoID, ShipName2ID("小型運輸艦"))
	assert.Equal(t, LargeCargoID, ShipName2ID("大型運輸艦"))
	assert.Equal(t, ColonyShipID, ShipName2ID("殖民船"))
	assert.Equal(t, RecyclerID, ShipName2ID("回收船"))
	assert.Equal(t, EspionageProbeID, ShipName2ID("間諜衛星"))
	assert.Equal(t, SolarSatelliteID, ShipName2ID("太陽能衛星"))
	assert.Equal(t, CrawlerID, ShipName2ID("履帶車"))
	assert.Equal(t, ReaperID, ShipName2ID("惡魔飛船"))
	assert.Equal(t, PathfinderID, ShipName2ID("探路者"))
}
