package com.caibeike.athena.service;

import com.caibeike.athena.model.Data;
import com.caibeike.athena.model.Domain;
import com.caibeike.athena.model.Experiment;
import com.caibeike.athena.model.Layer;

public interface AthenaService {

    String InsertDomain(Domain domain);

    String UpdateDomain(Domain domain);

    String DeleteDomain(Long domainId);

    String InsertLayerInDomain(Data data);

    String DeleteLayerInDomain(Data data);

    String UpdateLayer(Layer layer);

    String DeleteLayer(Long layerId);

    String InsertExperimentInLayer(Data data);

    String UpdateExperiment(Experiment experiment);

    String DeleteExperiment(Long expId);
}
