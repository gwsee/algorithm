package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const asinRunner = 10

type AwsFileCal struct {
	NationCode string
	one        sync.Once
	lock       sync.RWMutex
	filesCh    chan string
	Ids        map[string]uint
}

func main() {
	//var str = `1493, 1491, 5286, 5295, 1508, 1498, 1494, 5302, 5472, 5394, 5358, 1506, 5310, 1284, 1277, 5322, 1309, 5494, 1297, 1296, 1306, 5331, 1299, 5337, 5349, 5367, 1334, 1313, 5479, 4109, 5375, 5481, 5385, 5403, 5478, 5480, 5409, 5482, 5495, 5485, 5411, 5489, 5488, 5487, 5486, 5501, 5505, 5508, 5506, 1377, 5413, 5484, 5491, 5483, 5490, 5518, 5504, 5507, 1381, 5428, 5426, 5424, 5422, 5509, 5516, 5514, 1379, 5429, 5427, 5425, 5423, 5520, 5515, 5510, 5441, 5439, 5443, 5445, 5437, 5517, 1419, 5519, 1480, 4214, 1461, 1501, 1458, 1482, 1477, 5440, 5438, 5444, 5442, 1428, 5534, 1518, 1504, 5421, 1626, 1622, 1623, 5471, 1486, 5658, 4295, 5598, 4385, 5724, 4433, 8811, 8799, 5907, 5911, 4542, 4628, 8818, 8835, 8834, 8833, 8827, 8832, 8824, 8831, 8836, 8830, 8823, 8821, 8820, 8829, 8825, 8822, 8826, 8828, 8819, 8849, 8840, 8844, 8838, 8843, 8842, 8841, 8839, 8837, 4694, 8857, 8854, 8855, 8853, 8860, 8859, 8867, 8872, 8871, 8870, 8869, 8866, 8881, 8876, 8885, 8879, 8878, 8883, 8887, 8907, 8900, 8893, 8891, 8889, 8919, 4787, 8913, 4830, 4831, 4779, 3315, 8924, 3318, 4848, 3319, 4858, 4849, 4871, 4795, 3321, 3316, 4916, 4915, 4913, 4912, 4917, 4914, 4918, 4953, 4954, 4950, 4949, 5016, 5017, 5013, 5010, 8864, 5038, 8873, 8865, 8874, 8863, 8861, 5003, 5052, 8868, 8858, 4980, 8850, 8848, 8847, 8852, 8851, 8862, 8856, 8845, 5061, 4973, 8966, 5043, 5070, 8973, 8980, 8987, 5088, 5079, 5099, 5097, 5098, 5115, 5143, 5134, 5150, 5158, 5186, 5188, 5170, 5187, 5203, 5223, 5214, 5231, 3896, 5241, 5249, 5259, 5267, 5277, 4012, 11283, 1517, 1512, 1511, 1510, 1503, 1515, 1514, 1519, 1516, 1269, 1285, 1282, 1257, 1258, 1273, 1268, 1267, 1271, 1263, 1255, 1292, 1264, 1279, 1278, 1275, 1259, 1274, 1256, 1270, 1266, 1261, 1290, 1294, 1288, 1321, 1319, 1326, 1287, 1280, 1265, 1276, 1272, 1260, 1327, 1311, 1340, 1322, 1320, 1317, 1281, 1342, 1346, 1324, 1345, 1344, 1316, 1293, 1315, 1305, 1283, 1286, 1291, 1289, 1325, 1331, 1328, 1318, 1333, 1302, 1329, 1323, 1308, 1310, 1304, 1300, 1295, 1303, 1301, 1307, 1335, 1298, 1332, 1330, 1373, 1351, 1350, 1348, 1343, 1336, 1339, 1382, 1371, 1360, 1372, 1364, 1338, 1376, 1370, 1365, 1362, 1367, 1356, 1368, 1353, 1359, 1366, 1361, 1347, 1363, 1355, 1337, 1352, 1341, 1378, 1369, 1358, 1375, 1374, 1354, 1349, 1411, 1407, 1408, 1405, 1406, 1413, 1415, 1416, 1396, 1414, 1412, 1410, 1409, 1357, 1404, 1417, 1387, 1400, 1392, 1398, 1397, 1389, 1394, 1393, 1390, 1401, 1420, 1418, 1391, 1402, 1388, 1399, 1395, 1440, 1423, 1499, 1505, 1502, 1490, 1422, 1464, 1460, 1481, 1431, 1476, 1463, 1462, 1495, 1453, 1451, 1449, 1438, 1484, 1478, 1467, 1421, 1471, 1424, 1425, 1427, 1457, 1455, 1426, 1483, 1496, 1454, 1489, 1446, 1445, 1470, 1465, 1432, 1475, 1429, 1459, 1468, 1472, 1474, 1447, 1473, 1466, 1469, 1456, 1452, 1448, 1430, 1479, 1487, 1442, 1443, 1444, 1450, 1436, 1441, 1439, 1485, 1437, 1435, 1434, 1433, 1509, 1507, 1500, 1497, 1492, 1488, 1520, 1513, 1385, 1621, 1627, 1383, 1620, 1384, 1386, 1619, 1625, 1624, 1262, 1380, 1628, 3194, 3317, 3393, 3460, 3568, 3853, 3865, 8809, 8808, 8807, 8806, 8805, 8804, 8810, 8803, 8802, 8801, 8800, 8798, 11261, 11248, 11276, 11269`
	//arr:=strings.Split(str,",")
	//for _,v:=range arr {
	//	v = strings.TrimSpace(v)
	//	sql:=fmt.Sprintf("select count(1) from file_%v where same_mp = 'null' or same_mp is null;",v)
	//	fmt.Println(sql)
	//}
	one := new(AwsFileCal)
	one.Run()
}
func (a *AwsFileCal) Run() {
	a.one.Do(func() {
		a.Ids = make(map[string]uint)
		a.filesCh = make(chan string, asinRunner)
		for i := 0; i < asinRunner; i++ {
			go a.runner()
		}
	})
	var i uint
	var n uint
	for {
		time.Sleep(time.Second)
		if len(a.Ids) >= 10 {
			fmt.Println("-----------------------其中还有：", len(a.Ids), len(a.filesCh))
			time.Sleep(time.Second * 5)
			continue
		}
		fmt.Println("添加", n)
		for i = 0; i < 10; i++ {
			key := fmt.Sprintf("%v", i+n*100)
			a.Ids[key] = i
			a.filesCh <- key
		}
		n++
	}
}

func (a *AwsFileCal) runner() {
	for input := range a.filesCh {
		a.runOne(input)
	}
}
func (a *AwsFileCal) runOne(input string) {
	key := fmt.Sprintf("%v", input)
	defer func() {
		a.lock.Lock()
		delete(a.Ids, key)
		a.lock.Unlock()
	}()
	i := 1 + rand.Intn(10)
	fmt.Println("睡觉多少s:", i, key)
	time.Sleep(time.Second * time.Duration(i))
}